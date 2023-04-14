package main

import (
	"net/smtp"
	"os"
	"strconv"
)

func main() {
	// konfigurasi email server
	server := "smtp.gmail.com"
	port := 587
	email := os.Getenv("EMAIL")
	password := os.Getenv("PASSWORD")

	// konfigurasi email
	from := email
	to := []string{"jutionck@gmail.com", "fadli.rahman@enigmacamp.com"}
	subject := "Test Email"
	body := "This is a test email sent using Go."

	// menggabungkan konfigurasi email
	message := "From: " + from + "\n" +
		"To: " + to[0] + ", " + to[1] + "\n" +
		"Subject: " + subject + "\n\n" +
		body

	// mengirim email
	auth := smtp.PlainAuth("", email, password, server)
	err := smtp.SendMail(server+":"+strconv.Itoa(port), auth, email, to, []byte(message))
	if err != nil {
		panic(err)
	}
}
