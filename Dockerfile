FROM golang:alpine

WORKDIR /app
COPY . .
EXPOSE 8080

ADD id_rsa /root/.ssh/id_rsa
RUN touch /root/.ssh/known_hosts
RUN apk add -qU openssh
RUN ssh-keyscan github.com >> /root/.ssh/known_hosts

RUN go mod download
RUN go build -o main main.go
CMD ["./main"]