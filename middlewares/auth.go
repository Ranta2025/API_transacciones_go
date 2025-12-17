package middlewares

import (
	"database/sql"
	"net/http"
	"proyecto_transacciones/db"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/redis/go-redis/v9"
)
var secretkey = []byte("clave-secreta")

func Auth(data *sql.DB, rd *redis.Client) gin.HandlerFunc{
	return func(c *gin.Context){
		ctx := c.Request.Context()
		tokenstring := c.GetHeader("auth")

		if tokenstring == "" {
			c.JSON(http.StatusNonAuthoritativeInfo, gin.H{
				"acceso invalido": "token requerido",
			})
			c.Abort()
			return 
		}
		token, err := jwt.Parse(tokenstring, func(t *jwt.Token) (interface{}, error) {
			return secretkey, nil	
		})

		if err != nil || !token.Valid {
			c.JSON(http.StatusNonAuthoritativeInfo, gin.H{
				"acceso" : "token invalido",
			})
			c.Abort()
			return 
		}

		claims := token.Claims.(jwt.MapClaims)
		username := claims["username"].(string)
		rol := claims["rol"].(string)
		flag, _ := db.Check_user(ctx ,username ,rd, data)
		if !flag{
			c.JSON(http.StatusNonAuthoritativeInfo, gin.H{
				"acceso" : "token invalido",
			})
			c.Abort()
			return 
		}

		c.Set("username", username)
		c.Set("rol", rol)
		c.Next()
	}
}