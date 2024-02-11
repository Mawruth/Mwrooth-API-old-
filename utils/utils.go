package utils

import (
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"io"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func GenerateToken(email string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": email,
		"role":  "USER",
		"exp":   time.Now().Add(time.Hour * 24 * 30).Unix(),
	},
	)

	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET_KEY")))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func GenerateVerificationMessage(code string) []byte {
	msgFile, err := os.Open("templates/verification_msg.txt")
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}
	defer msgFile.Close()
	msg, err := io.ReadAll(msgFile)
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}
	return []byte(fmt.Sprintf("%s\n\n\n%s", string(msg), code))
}

func GenerateOTP() string {
	code := rand.Intn(10000)
	return fmt.Sprintf("%d", code)
}

func ParseOTP(otp string) (string, time.Time) {
	timestamp := strings.Split(otp, "|")[2]
	expTime, err := strconv.ParseInt(strings.Split(otp, "|")[1], 10, 32)
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}

	rawOTP := strings.Split(otp, "|")[0]
	timestampAsTime, err := time.Parse("2006-01-02 15:04:05", timestamp)
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}
	timestampAsTime = timestampAsTime.Add(time.Minute * time.Duration(expTime))
	return rawOTP, timestampAsTime
}
