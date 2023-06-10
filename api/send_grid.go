package api

import (
	"fmt"
	"log"
	"os"

	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

type EmailModel struct {
	Username string `json:"username"`
	Email    string `json:"email"`
}

func SendEmail(params []EmailModel) error {

	form := mail.NewEmail(os.Getenv("FROM_USER"), os.Getenv("FROM_ADDRESS"))
	subject := "TEST MAIL"

	message := mail.NewV3Mail()
	message.SetFrom(form)
	message.SetTemplateID(os.Getenv("SEND_GRID_TEMPLATE_ID"))
	message.Subject = subject

	for _, param := range params {
		p := mail.NewPersonalization()
		p.AddTos(mail.NewEmail(param.Username, param.Email))
		p.SetDynamicTemplateData("username", param.Username)
		message.AddPersonalizations(p)
	}

	client := sendgrid.NewSendClient(os.Getenv("SEND_GRID_API_KEY"))
	response, err := client.Send(message)
	if err != nil {
		log.Printf("Error sending mail: %v", err)
		return err
	}
	fmt.Println(response.Body)
	fmt.Println(response.StatusCode)
	log.Print("Success sending mail")
	return nil
}
