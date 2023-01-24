package controllers

import (
	"net/http"

	"app/models"

	"github.com/gin-gonic/gin"
	"gopkg.in/olahol/melody.v1"
	"github.com/dgrijalva/jwt-go"
)

func WebSocket(c *gin.Context) {

	m := melody.New()

	h := models.TestHeader{}

	if err := c.ShouldBindHeader(&h); err != nil {
		c.JSON(200, err)
	}
	
	cookie, err := c.Cookie(h.Name)

	if err != nil {

		if err == http.ErrNoCookie {
			c.JSON(http.StatusUnauthorized, "")
			return
		}

		c.JSON(http.StatusBadRequest, "")
		return
	}

	claims := &models.Claims{}

	tkn, err := jwt.ParseWithClaims(cookie, claims,
	func(*jwt.Token) (interface{}, error) {
		return []byte(h.Name), nil
	})

	if err != nil {

		if err == jwt.ErrSignatureInvalid {
			c.JSON(http.StatusUnauthorized, "")
			return
		}

		c.JSON(http.StatusBadRequest, "")
		return
	}

	if !tkn.Valid {
		c.JSON(http.StatusUnauthorized, "")
		return
	}

	m.HandleRequest(c.Writer, c.Request)

	c.String(200, "Hello %s", h.Name)
}