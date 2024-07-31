package email

import (
	"log"
	"net/smtp"
)

func Send(address string) {
	// Choose authentication method and set it up
	auth := smtp.PlainAuth("", "john.doe@gmail.com", "extremely_secret_pass", "smtp.gmail.com")

	// Connect to the server, set up a message, and send it
	err := smtp.SendMail("smtp.gmail.com:587", auth, "john.doe@gmail.com", []string{address}, []byte("Subject: Hello!\n\nThis is the email body."))
	if err != nil {
		log.Fatal(err)
	}
}
