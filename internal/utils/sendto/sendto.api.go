package sendto

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type MailRequest struct {
	ToEmail     string `json:"toEmail"`
	MessageBody string `json:"messageBody"`
	Subject     string `json:"subject"`
	Attachment  string `json:"attachment"`
}

func SendEmailToJavaByAPI(otp string, email string, purpose string) error {
	// URL API
	postURL := "http://localhost:8080/email/send_text"

	// Data Json
	mailRequest := MailRequest{
		ToEmail:     email,
		MessageBody: "OTP IS " + otp,
		Subject:     "Verify OTP " + purpose,
		Attachment:  "path/to/email",
	}

	// convert struc to json
	requestBobdy, err := json.Marshal(mailRequest)
	if err != nil {
		return err
	}
	// create request
	req, err := http.NewRequest("POST", postURL, bytes.NewBuffer(requestBobdy))
	if err != nil {
		return err
	}
	// PUT header
	req.Header.Set("Content-Type", "application/json")

	// execute request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Println("Error while closing response body", err)
		}
	}(resp.Body)
	fmt.Sprintln("Response status: ", resp.Status)
	return nil
}
