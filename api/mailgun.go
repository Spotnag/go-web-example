package api

import (
	"context"
	"fmt"
	"github.com/mailgun/mailgun-go/v4"
	"go-web-example/config"
	"time"
)

type MailgunService struct {
	mailgun *mailgun.MailgunImpl
}

func NewMailgunService() (*MailgunService, error) {
	// Your available domain names can be found here:
	// (https://app.mailgun.com/app/domains)
	var yourDomain = config.AppConfig.MailgunDomain // e.g. mg.yourcompany.com

	// You can find the Private API Key in your Account Menu, under "Settings":
	// (https://app.mailgun.com/app/account/security)
	var privateAPIKey = config.AppConfig.MailgunAPIKey
	// Create an instance of the Mailgun Client
	mg := mailgun.NewMailgun(yourDomain, privateAPIKey)

	//When you have an EU-domain, you must specify the endpoint:
	//mg.SetAPIBase("https://api.eu.mailgun.net/v3")

	return &MailgunService{
		mailgun: mg,
	}, nil

}

func (mg *MailgunService) SendEmail() error {
	sender := "sender@example.com"
	subject := "Fancy subject!"
	body := "Hello from Mailgun Go!"
	recipient := "martyn.north@hotmail.co.uk"

	// The message object allows you to add attachments and Bcc recipients
	message := mg.mailgun.NewMessage(sender, subject, body, recipient)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	// Send the message with a 10 second timeout
	resp, id, err := mg.mailgun.Send(ctx, message)

	if err != nil {
		return err
	}

	fmt.Printf("ID: %s Resp: %s\n", id, resp)
	return nil
}
