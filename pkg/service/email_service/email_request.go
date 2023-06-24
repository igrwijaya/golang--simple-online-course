package email_service

type EmailVerificationRequest struct {
	Subject          string
	Email            string
	VerificationCode string
}

type ForgotPasswordRequest struct {
	Subject string
	Email   string
	Code    string
}
