package controllers

import (
	"encoding/json"
	"github.com/amirjavadi/go_auth_api/pkg/config"
	"github.com/amirjavadi/go_auth_api/pkg/models"
	"github.com/amirjavadi/go_auth_api/pkg/utils"
	"github.com/amirjavadi/go_auth_api/pkg/validators"
	"net/http"
)

func Register(w http.ResponseWriter, r *http.Request) {
	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "request is invalid", http.StatusBadRequest)
		return
	}
	if err := validators.AuthValidation(&user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	hashedPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		http.Error(w, "خطا در هش کردن پسورد", http.StatusBadRequest)
	}

	user.Password = hashedPassword
	result := config.GetDb().Create(&user)

	if result.Error != nil {
		http.Error(w, "خطا در ثبت نام کاربر", http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "ثبت نام با موفقیت انجام شد"})
}

func Login(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "invalid request", http.StatusBadRequest)
	}

	var user models.User
	result := config.GetDb().Where("email=?", input.Email).First(&user)
	if result.Error != nil {
		http.Error(w, "کاربر یافت نشد", http.StatusUnauthorized)
	}
	err := utils.CheckPassword(user.Password, input.Password)
	if err != nil {
		http.Error(w, "password is incorrect", http.StatusUnauthorized)
	}

	token, er := utils.GenerateJwt(user.ID)
	if er != nil {
		http.Error(w, "خطا در تولید توکن", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(map[string]string{"token": token})
}

func Logout(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(map[string]string{"message": "با موفقیت خارج شدید"})
	// you should handle logout in frontend
}
