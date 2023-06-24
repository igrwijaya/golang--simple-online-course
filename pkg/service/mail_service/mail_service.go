package mail_service

import (
	"bytes"
	"errors"
	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
	"golang-online-course/pkg/response"
	"html/template"
	"os"
	"path/filepath"
)

type Service interface {
	SendVerification(requestDto EmailVerificationDto) response.Basic
}

type mailService struct {
}

func (service *mailService) SendVerification(requestDto EmailVerificationDto) response.Basic {
	appDir, _ := os.Getwd()
	templateFilePath := filepath.Join(appDir, "/web/template/verification_email.html")

	htmlContent, errParsingTemplate := parseTemplate(templateFilePath, requestDto)

	if errParsingTemplate != nil {
		return response.Basic{
			Code:  400,
			Error: errParsingTemplate,
		}
	}

	return sendEmail(requestDto.Email, requestDto.Subject, htmlContent)
}

func NewService() Service {
	return &mailService{}
}

//#region Private Func

func sendEmail(emailReceiver string, subject string, htmlContent string) response.Basic {
	emailSender := os.Getenv("MAIL_SENDER_NAME")
	sendgridKey := os.Getenv("MAIL_KEY")

	from := mail.NewEmail(emailSender, emailSender)
	to := mail.NewEmail(emailReceiver, emailReceiver)
	message := mail.NewSingleEmail(from, subject, to, "", htmlContent)

	sendGridClient := sendgrid.NewSendClient(sendgridKey)

	sendResponse, errSendEmail := sendGridClient.Send(message)

	if errSendEmail != nil {
		return response.Basic{
			Code:  400,
			Error: errSendEmail,
		}
	} else if sendResponse.StatusCode == 200 || sendResponse.StatusCode == 202 {
		return response.Success()
	}

	return response.Basic{
		Code:  400,
		Error: errors.New("sendgrid not returned success response. " + sendResponse.Body),
	}
}

func parseTemplate(templateFilePath string, data interface{}) (string, error) {
	templateContent, errReadTemplate := template.ParseFiles(templateFilePath)

	if errReadTemplate != nil {
		return "", errReadTemplate
	}

	contentWriter := new(bytes.Buffer)

	errWriteContent := templateContent.Execute(contentWriter, data)

	if errWriteContent != nil {
		return "", errWriteContent
	}

	return contentWriter.String(), nil
}

//#endregion
