package validators

import (
	"errors"
	"github.com/amirjavadi/go_auth_api/pkg/models"
	"regexp"
	"unicode"
)

func AuthValidation(user *models.User) error {

	if user.Username == "" {
		return errors.New("نام کاربری نمیتواند خالی باشد")
	}
	if !emailValidation(user.Email) {
		return errors.New("قرمت ایمیل نادرست است")
	}
	if !mobileValidation(user.Mobile) {
		return errors.New("قرمت موبایل نادرست است")
	}
	if len(user.Password) < 6 || !containLetterAndNumber(user.Password) {
		return errors.New("رمز عبور باید حداقل ۸ کاراکتر و شامل حرف و عدد باشد")
	}

	return nil
}

func emailValidation(email string) bool {
	regex := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,}$`)
	return regex.MatchString(email)
}

func mobileValidation(mobile string) bool {
	regex := regexp.MustCompile(`^09[0-9]{9}$`)
	return regex.MatchString(mobile)
}

func containLetterAndNumber(s string) bool {
	hasNum, hasLet := false, false

	for _, value := range s {
		if unicode.IsLetter(value) {
			hasLet = true
		}
		if unicode.IsNumber(value) {
			hasNum = true
		}
	}
	return hasNum && hasLet
}
