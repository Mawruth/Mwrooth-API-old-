package utils

import (
	"bytes"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/golang-jwt/jwt/v5"
	"github.com/h2non/filetype"
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

func UploadImageToS3(file io.Reader) (string, error) {
	allowedExtensions := []string{"jpg", "jpeg", "png"}

	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(os.Getenv("S3_REGION")),
	})
	if err != nil {
		return "", err
	}
	fmt.Println(sess.Config.Credentials.Get())

	buf, err := io.ReadAll(file)
	if err != nil {
		return "", err
	}

	if len(buf) > 1e7 {
		return "", fmt.Errorf("File size is exceeds 2MB limit")
	}

	kind, _ := filetype.Match(buf)
	if !contains(allowedExtensions, kind.Extension) {
		return "", fmt.Errorf("Unsupported file type. please upload jpg, jpeg or png")
	}

	uploader := s3manager.NewUploader(sess)
	extension := kind.Extension
	fileName := generateUniqueTimestamp() + "." + extension

	upParams := &s3manager.UploadInput{
		Bucket:      aws.String(os.Getenv("S3_BUCKET")),
		Key:         aws.String(fmt.Sprintf("%s", fileName)),
		Body:        bytes.NewReader(buf),
		ContentType: aws.String(kind.MIME.Value),
		ACL:         aws.String("public-read"),
	}

	result, err := uploader.Upload(upParams)
	if err != nil {
		return "", fmt.Errorf("failed to upload file, %v", err)
	}

	return result.Location, nil
}

func generateUniqueTimestamp() string {
	return fmt.Sprintf("%d", time.Now().UnixNano())
}

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
