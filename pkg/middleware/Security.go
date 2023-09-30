package middleware

import (
	"errors"
	"net/http"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/nicoxxg/go-server/internal/domain/cliente"
)

// IDEA crear un metodo que le pase de parametro el token y me retorne un clente

func tokenToClient(tokenString string) (cliente.ClientJson, error) {

	secretKey := []byte(os.Getenv("TOKEN"))
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Verifica el método de firma
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("Método de firma no válido")
		}
		// Define y retorna tu clave secreta aquí
		return secretKey, nil
	})
	if err != nil {
		return cliente.ClientJson{}, err
	}
	claims := token.Claims.(jwt.MapClaims)

	var cliente cliente.ClientJson

	cliente.Email = claims["email"].(string)

	cliente.Password = claims["password"].(string)

	cliente.Authority = claims["rol"].(string)

	cliente.Exp = int64(claims["exp"].(float64))

	return cliente, nil
}

func RolVerification(rolParametro string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		tokenString := ctx.GetHeader("Authorization")

		token, err := tokenToClient(tokenString)

		if err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"message": "Token de autorización inválido: " + err.Error(),
			})
			ctx.Abort()
			return
		}

		rolUsuario := token.Authority
		if rolUsuario != rolParametro {
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
		secretKey := []byte(os.Getenv("TOKEN"))

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			// Verifica el método de firma
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, errors.New("Método de firma no válido")
			}
			// Define y retorna tu clave secreta aquí
			return secretKey, nil
		})

		if err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"message": "Token de autorización inválido: " + err.Error(),
			})
			ctx.Abort()
			return
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			exp := int64(claims["exp"].(float64))
			now := time.Now().Unix()

			if now > exp {

				ctx.JSON(http.StatusUnauthorized, gin.H{
					"message": "Token de autorización ah expirado",
				})
				ctx.Abort()
				return
			}
			newExp := now + (2 * 3600)

			claims["exp"] = newExp

			newToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

			newTokenString, err := newToken.SignedString(secretKey)

			if err != nil {
				ctx.JSON(http.StatusUnauthorized, gin.H{
					"message": "Error al firmar el nuevo token: " + err.Error(),
				})
				ctx.Abort()
				return
			}
			// Devuelve el nuevo token en la respuesta
			ctx.Header("Authorization", "Bearer "+newTokenString)

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
