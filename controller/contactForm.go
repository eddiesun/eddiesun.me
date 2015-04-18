package controller

import (
	"github.com/eddiesun.me/config"
	"github.com/sendgrid/sendgrid-go"
	"log"
	"net/http"
)

func ContactForm(w http.ResponseWriter, r *http.Request) {
	log.Println("ContactFormController - incoming request")

	// place holder param string
	var messageSent int = 0
	var errorMessage string

	// get form variables
	var name string = r.PostFormValue("name")
	var email string = r.PostFormValue("email")
	var message string = r.PostFormValue("message")

	// send an email to my email address.
	sg := sendgrid.NewSendGridClient(config.SENDGRID_USERNAME, config.SENDGRID_PASSWORD)
	mail := sendgrid.NewMail()
	mail.AddTo(config.MY_EMAIL_ADDRESS)
	mail.AddToName(config.MY_NAME)
	mail.SetSubject("New message from eddiesun.me")
	mail.SetHTML(name + " <" + email + "> sent you the following message:<br/>------------------------------------------<br/><pre>" + message + "</pre>")
	mail.SetFrom(email)
	mail.SetFromName(name)
	if r := sg.Send(mail); r == nil {
		log.Println("Contact me message sent")
		messageSent = 1
	} else {
		log.Println("Contact me message ERROR")
		log.Printf("%+v\n", r)
		messageSent = 0
		errorMessage = "Sorry, we failed to send your message..."
	}

	// redirect user
	http.Redirect(w, r, "/?messageSent="+string(messageSent)+"&errorMessage="+errorMessage+"#contact", http.StatusFound)
}
