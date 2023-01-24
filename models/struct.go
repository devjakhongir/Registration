package models

import "github.com/dgrijalva/jwt-go"

type Mail struct {
	Email string `json:"email"`
}

type Login struct{
	Username string `gorm:"username" json:"username"`
	Email 	 string `gorm:"email" 	 json:"email"`
	Password string `gorm:"password" json:"password"`
}

type Token struct {
	AccessToken string  `gorm:"auth"`
	UserLogin 	Login   `gorm:"-"`
}

type ChangePassword struct {
	Email 	 		string `json:"email"`
	CurrentPassword string `json:"currentPassword"`
	NewPassword 	string `json:"newPassword"`
}

type Claims struct {
	jwt.StandardClaims
}

type TestHeader struct {
	Name   	string	`header:"name"`
}