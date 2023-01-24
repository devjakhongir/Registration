package controllers

import (
	"io"
	"log"
	"net/http"
	"crypto/sha1"
	"encoding/hex"

	"app/utils"
	"app/models"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {

	login := models.Login{}

	c.BindJSON(&login)

	h := sha1.New()

	io.WriteString(h, login.Email)

	log.Println(hex.EncodeToString(h.Sum(nil)))

	if hex.EncodeToString(h.Sum(nil)) == login.Password {

		auth := models.Token{}

		tokenPassword, err := utils.CreateJWT(login.Username)

		if err != nil {
			c.JSON(http.StatusInternalServerError, models.Login{})
			return
		}

		auth.AccessToken = tokenPassword
		auth.UserLogin = login

		db, err := utils.DB()

		if err != nil { 
			log.Print(err) 
			c.JSON(http.StatusInternalServerError, models.Login{})
			return
		}

		rows := db.Table("registration").
			Select("email, password").
			Where("password = ?", login.Password).
			Find(&login)

		if rows.RowsAffected == 1{
			c.String(http.StatusBadRequest, "You are on the list : %v", login.Email)
			return
		}

		result := db.Table("registration").Create(&login)

		if utils.IsNotFound(result) {
			log.Println(result.Error)
			return
		}

		cookie, err := c.Cookie(login.Username)

		if err != nil {
			cookie = "NotSet"
			c.SetCookie(login.Username, auth.AccessToken, 3600, "/", "localhost", false, true)
		}
	
		log.Printf("Cookie value: %s \n", cookie)
		
		c.String(http.StatusCreated, "%v", auth)
	} else {

		c.JSON(http.StatusUnauthorized, models.Login{})
	}
}