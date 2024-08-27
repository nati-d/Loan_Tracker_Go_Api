package userusecase

import (
	"loan_tracker/bootstrap"
	"loan_tracker/config"
	"loan_tracker/domain"
)


func (uc *UserUsecase) PasswordResetRequest(email string) error {
	user, err := uc.UserRepository.GetUserByUsernameOrEmail(email)

	if err != nil {
		return err
	}

	
	apiBase,err:= bootstrap.GetEnv("API_BASE")
	if err != nil {
		return err
	}
	resetClaims := &domain.PasswordResetClaims{
		Username: user.Username,
	}
	resteToken, err := config.GenerateToken(resetClaims)
	if err != nil {
		return err
	}

	emailHeader := "Welcome to Loan Tracker - Reset Your Password"
	emailBody := `
		<!DOCTYPE html>
		<html>
		<head>
			<title>Reset Your Password</title>
		</head>
		<body style="font-family: Arial, sans-serif; line-height: 1.6; color: #333;">
			<h2 style="color: #4CAF50;">Welcome to Blogs, ` + user.Username + `!</h2>
			<p>Thank you for registering with us. To complete your registration, please Reset Your Password address by clicking the button below:</p>
			<p style="text-align: center;">
				<a href="` + apiBase + `/users/verify-email?token=` + resteToken + `" style="display: inline-block; padding: 10px 20px; margin: 20px 0; font-size: 18px; color: #ffffff; background-color: #4CAF50; text-decoration: none; border-radius: 5px;">Verify Email</a>
			</p>
			<p>If you did not create an account with us, please ignore this email.</p>
			<p>Best regards,<br>The Loaner Team</p>
			<hr style="border: none; border-top: 1px solid #eee; margin: 20px 0;">
			<p style="font-size: 12px; color: #999;">If you have any questions or need assistance, please contact our support team at <a href="mailto:support@blogs.com">support@loaner.com</a>.</p>
		</body>
		</html>
	`

	err = config.SendEmail(user.Email, emailHeader, emailBody, true)
	if err != nil {
		return err
	}

	return nil




}