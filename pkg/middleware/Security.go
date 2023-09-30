package middleware

import (
	"errors"
	"net/http"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func RolVerification(rolParametro string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		tokenString := ctx.GetHeader("Authorization")
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			// Verifica el método de firma
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, errors.New("Método de firma no válido")
			}
			// Define y retorna tu clave secreta aquí
			return os.Getenv("TOKEN"), nil
		})
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"message": "Token de autorización inválido",
			})
			ctx.Abort()
			return
		}

		claims := token.Claims.(jwt.MapClaims)

		rolUsuario := claims["rol"].(string)
		if rolUsuario != rolUsuario {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"message": "Usuario No authorizado",
			})
			ctx.Abort()
			return
		}

		ctx.Next()
	}
}

func Verification() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		tokenString := ctx.GetHeader("Authorization")
		if tokenString == "" {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"message": "Token de autorizacion faltante",
			})
			ctx.Abort()
			return
		}
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			// Verifica el método de firma
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, errors.New("Método de firma no válido")
			}
			// Define y retorna tu clave secreta aquí
			return os.Getenv("TOKEN"), nil
		})

		if err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"message": "Token de autorización inválido",
			})
			ctx.Abort()
			return
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			exp := int64(claims["exp"].(float64))

			if time.Now().Unix() > exp {
				ctx.JSON(http.StatusUnauthorized, gin.H{
					"message": "Token de autorización ah expirado",
				})
				ctx.Abort()
				return
			}

			ctx.Next()
		} else {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"message": "Token de autorización inválido",
			})
			ctx.Abort()
			return
		}

	}
}
