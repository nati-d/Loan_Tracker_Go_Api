package config

import (
	"errors"
	"regexp"
)



func ValidateUsername(username string) error {
	// Define the allowed characters using a regular expression
	const usernamePattern = `^[a-zA-Z0-9_-]+$`
	matched, err := regexp.MatchString(usernamePattern, username)
	if err != nil {
		return err
	}
	if !matched {
		return errors.New("Enter Valid Username")
	}
	return nil
}

func ValidatePassword(password string) error {
	if len(password) < 8 {
		return errors.New("password must be at least 8 characters long")
	}

	// Define regular expressions for the different character types
	var (
		uppercasePattern  = `[A-Z]`
		lowercasePattern  = `[a-z]`
		digitPattern      = `[0-9]`
		specialCharPattern = `[!@#\$%\^&\*\(\)_\+\-=\[\]\{\};:'"\|\\,.<>\/?]`
	)

	// Check for at least one uppercase letter
	if matched, _ := regexp.MatchString(uppercasePattern, password); !matched {
		return errors.New("password must contain at least one uppercase letter")
	}

	// Check for at least one lowercase letter
	if matched, _ := regexp.MatchString(lowercasePattern, password); !matched {
		return errors.New("password must contain at least one lowercase letter")
	}

	// Check for at least one digit
	if matched, _ := regexp.MatchString(digitPattern, password); !matched {
		return errors.New("password must contain at least one digit")
	}

	// Check for at least one special character
	if matched, _ := regexp.MatchString(specialCharPattern, password); !matched {
		return errors.New("password must contain at least one special character")
	}

	return nil
}

func ValidateEmail(email string) error {
	// Define a regular expression for a simple email validation
	var emailRegex = `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`

	// Compile the regular expression
	re := regexp.MustCompile(emailRegex)

	// Validate the email using the regular expression
	if !re.MatchString(email) {
		return errors.New("invalid email format")
	}

	return nil
}

