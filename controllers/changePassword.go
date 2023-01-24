package controllers

import (
	"log"
	"net/http"

	"app/utils"
	"app/models"

	"github.com/gin-gonic/gin"
)

func PostChangePassword(c *gin.Context) {

	change := models.ChangePassword{}

	c.BindJSON(&change)

	db, err := utils.DB()

	if err != nil { 
		log.Print(err) 
		c.JSON(http.StatusInternalServerError, models.Login{})
		return
	}

	login := models.Login{}

	rows := db.Table("registration").
		Select("email, password").
		Where("password = ? and email = ?", change.CurrentPassword, change.Email).
		Find(&login)

	log.Println(login)

	if rows.RowsAffected != 1{
		c.String(http.StatusBadRequest, "Error Password or Email\n%s\n%s", change.Email, change.CurrentPassword)
		return
	}

	rowsUpdate := db.Table("registration").
		Where("email = ?", login.Email).
		Update("password", change.NewPassword)

	if utils.IsNotFound(rowsUpdate) {
		log.Println(rowsUpdate.Error)
		return
	}

	auth := models.Token{}

	tokenPassword, err := utils.CreateJWT(change.Email)

	if err != nil {
		c.JSON(http.StatusInternalServerError, models.Login{})
		return
	}

	login.Password = change.NewPassword
	auth.AccessToken = tokenPassword
	auth.UserLogin = login

	cookie, _ := c.Cookie(change.Email)

	c.SetCookie(change.Email, auth.AccessToken, 3600, "/", "localhost", false, true)

	log.Printf("Cookie value: %s \n", cookie)

	log.Println("OK")

	c.String(http.StatusOK, "%v", change)
}