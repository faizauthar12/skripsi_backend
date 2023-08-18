package email

import (
	"fmt"
	"net/smtp"
	"testing"

	"github.com/joho/godotenv"
)

func TestSendOrderConfirmation(t *testing.T) {

	errorLoadEnv := godotenv.Load()
	if errorLoadEnv != nil {
		t.Log(errorLoadEnv)
		t.FailNow()
	}

	configDotEnv := ConfigDotEnvv()

	auth := LoginAuth(configDotEnv.EmailSender, configDotEnv.PasswordSender)
	subject := fmt.Sprintf("Subject: Order for blablabla")
	message := fmt.Sprintf("Rincian pemesanan bla bla bla")

	// Here we do it all: connect to our server, set up a message and send it
	to := []string{configDotEnv.EmailReceier2}
	msg := []byte(subject +
		"\r\n" +
		message)

	err := smtp.SendMail("smtp.gmail.com:587", auth, configDotEnv.EmailSender, to, msg)

	if err != nil {
		t.Log(err)
		t.FailNow()
	}
}
