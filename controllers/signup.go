package controllers

import (
	"io"
	"log"
	"net/http"
	"crypto/sha1"
	"encoding/hex"

	"app/models"
	"app/utils"

	"github.com/gin-gonic/gin"
)


func SignUp(c *gin.Context) {

	var mail models.Mail
	
	c.BindJSON(&mail)

	h := sha1.New()

	io.WriteString(h, mail.Email)

	passwordSHa1 := hex.EncodeToString(h.Sum(nil))

	db, err := utils.DB()

	if err != nil { 
		log.Print(err) 
		c.JSON(http.StatusInternalServerError, models.Login{})
	}

	var login models.Login

	rows := db.Table("registration").
		Select("email, password").
		Where("email = ?", mail.Email).
		Find(&login)

	if rows.RowsAffected == 1{
		c.String(http.StatusBadRequest, "You are in the registration section : %v", login.Email)
	} else {
		utils.Send(passwordSHa1, mail.Email)
		c.String(http.StatusOK, "Send %s email message...\n%d", mail, passwordSHa1)
	}
}