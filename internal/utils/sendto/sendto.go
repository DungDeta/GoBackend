package sendto

import (
	"fmt"
	"net/smtp"
	"strings"

	"go.uber.org/zap"
	"myproject/global"
)

const (
	// SMTPServer is the SMTP server address
	SMTPServer = "smtp.gmail.com"
	// SMTPPort is the SMTP server port
	SMTPPort = "587"
	SMTPUser = "macbadao@gmail.com"
	SMTPPass = "Duydung20041998"
)

type EmailAddress struct {
	Address string `json:"address"`
	Name    string `json:"name"`
}
type Mail struct {
	From    EmailAddress
	To      []string
	Subject string
	Body    string
}

func BuildMail(mail Mail) string {
	msg := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n"
	msg += fmt.Sprintf("From: %s\r\n", mail.From.Address)
	msg += fmt.Sprintf("To: %s\r\n", strings.Join(mail.To, ";"))
	msg += fmt.Sprintf("Subject: %s\r\n", mail.Subject)
	msg += fmt.Sprintf("\r\n%s\r\n", mail.Body)

	return msg
}
func SendTextEmail(to []string, from string, otp int) error {
	contentMail := Mail{
		From: EmailAddress{
			Address: from,
			Name:    "MyProject",
		},
		To:      to,
		Subject: "OTP",
		Body:    "Your OTP is " + string(rune(otp)),
	}
	msg := BuildMail(contentMail)
	// Send using SMTP
	fmt.Println(msg)
	auth := smtp.PlainAuth("", SMTPUser, SMTPPass, SMTPServer)
	err := smtp.SendMail(SMTPServer+":"+SMTPPort, auth, SMTPUser, to, []byte(msg))
	if err != nil {
		global.Logger.Error("Send mail error", zap.Error(err))
		return err
	}

	return nil
}
