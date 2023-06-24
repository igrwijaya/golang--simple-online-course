package mail_service

type EmailVerificationDto struct {
	Subject          string
	Email            string
	VerificationCode string
}
