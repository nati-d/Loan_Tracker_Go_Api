package userusecase

import (
	"errors"
	"loan_tracker/bootstrap"
	"loan_tracker/config"
	"loan_tracker/domain"
)

func (u *UserUsecase) RegisterUser(user *domain.User) error {

	// Check if username or email already exists
	val, err := u.UserRepository.CheckUserExists(user.Username)
	if err != nil {
		return err
	}
	if val {
		return errors.New("Username already exists")
	}

	val, err = u.UserRepository.CheckUserExists(user.Email)
	if err != nil {
		return err
	}
	if val {
		return errors.New("Email already exists")
	}

	// Validate the username, email, and password
	err = config.ValidateUsername(user.Username)
	if err != nil {
		return err
	}
	err = config.ValidateEmail(user.Email)
	if err != nil {
		return err
	}
	err = config.ValidatePassword(user.Password)
	if err != nil {
		return err
	}

	// Hash the password
	hashedPassword, err := config.HashPassword(user.Password)
	if err != nil {
		return err
	}
	user.Password = hashedPassword

	// Get the API base URL from environment variables
	apiBase, err := bootstrap.GetEnv("API_BASE")
	if err != nil {
		return err
	}

	// Generate a verification token
	verifyToken, err := config.GenerateToken(&domain.RegisterClaims{User: *user})
	if err != nil {
		return err
	}

	// Create a professional email body for verification
	emailHeader := "Welcome to Loan Tracker - Verify Your Email"
	emailBody := `
		<!DOCTYPE html>
		<html>
		<head>
			<title>Verify Your Email</title>
		</head>
		<body style="font-family: Arial, sans-serif; line-height: 1.6; color: #333;">
			<h2 style="color: #4CAF50;">Welcome to Blogs, ` + user.Username + `!</h2>
			<p>Thank you for registering with us. To complete your registration, please verify your email address by clicking the button below:</p>
			<p style="text-align: center;">
				<a href="` + apiBase + `/users/verify-email?token=` + verifyToken + `" style="display: inline-block; padding: 10px 20px; margin: 20px 0; font-size: 18px; color: #ffffff; background-color: #4CAF50; text-decoration: none; border-radius: 5px;">Verify Email</a>
			</p>
			<p>If you did not create an account with us, please ignore this email.</p>
			<p>Best regards,<br>The Loaner Team</p>
			<hr style="border: none; border-top: 1px solid #eee; margin: 20px 0;">
			<p style="font-size: 12px; color: #999;">If you have any questions or need assistance, please contact our support team at <a href="mailto:support@blogs.com">support@loaner.com</a>.</p>
		</body>
		</html>
	`

	// Send the verification email
	err = config.SendEmail(user.Email, emailHeader, emailBody, true)
	if err != nil {
		return err
	}

	return nil
}
