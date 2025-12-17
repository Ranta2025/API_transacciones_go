package middlewares

import (
	"net/http"
	"proyecto_transacciones/utils"
	"strings"
	"github.com/gin-gonic/gin"
)

func Validation_param_into(c *gin.Context)  {
	c.Request.Context()
	username := c.PostForm("username")
	if username == "" {
		c.String(http.StatusBadRequest, "campo usuario obligatorio")
		c.Abort()
		return 
	}

	gmail := c.PostForm("gmail")
	if gmail == "" {
		c.String(http.StatusBadRequest, "campo gmail obligatorio")
		c.Abort()
		return 
	}
	
	password := c.PostForm("password")
	if password == "" {
		c.String(http.StatusBadRequest, "campo password obligatorio")
		c.Abort()
		return 
	}
	c.Next()
}

func Validation_param_len(c *gin.Context) {
	username := c.PostForm("username")
	if len(username) < 5 {
		c.String(http.StatusBadRequest, "El usuario debe contener mas de cinco caracteres")
		c.Abort()
		return 
	}

	gmail := c.PostForm("gmail")
	if !strings.HasSuffix(gmail, "@gmail.com") {
		c.String(http.StatusBadRequest, "El gmail tiene que terminar con @gmail.com")
		c.Abort()
		return 
	}
	if len(gmail) < 11 {
		c.String(http.StatusBadRequest, "El gmail debe contener una cadena de texto aparte de la estandar @gmail.com")
		c.Abort()
		return 
	}

	password := c.PostForm("password")
	if len(password) < 8 {
		c.String(http.StatusBadRequest, "La contraseña tiene que contener mas de 8 caracteres")
		c.Abort()
		return 
	}
	c.Next()
}

func Validation_password(c *gin.Context){
	password := c.PostForm("password")
	if !utils.Validation_Letter(password){
		c.String(http.StatusBadRequest, "La contraseña debe contener al menos una letra")
		c.Abort()
		return 
	}

	if !utils.Validation_Number(password){
		c.String(http.StatusBadRequest, "La contraseña debe contener al menos un numero")
		c.Abort()
		return 
	}

	if !utils.Validation_Character_special(password){
		c.String(http.StatusBadRequest, "La contraseña debe contener al menos un caracter especial")
		c.Abort()
		return 
	}

	if !utils.Validation_Upper(password){
		c.String(http.StatusBadRequest, "La contraseña debe contener al menos una letra Mayuscula")
		c.Abort()
		return 
	}
	c.Next()
}