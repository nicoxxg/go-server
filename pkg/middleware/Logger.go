package middleware

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/nicoxxg/go-server/internal/domain/cliente"
	"golang.org/x/crypto/bcrypt"
)

type security struct {
	ClienteRepository cliente.ClienteRepository
}

type Security interface {
	Logger() gin.HandlerFunc
}

func NewSecurity(repository cliente.ClienteRepository) Security {
	return &security{
		ClienteRepository: repository,
	}
}

func (s *security) Logger() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		err := godotenv.Load()
		if err != nil {
			log.Fatal(err)
		}

		email := ctx.Query("email")

		password := ctx.Query("password")

		clienteObteido, err := s.ClienteRepository.FindByEmail(ctx, email)

		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"message": "Cliente no encontrado",
			})
			return
		}

		err = bcrypt.CompareHashAndPassword([]byte(clienteObteido.Contrasena), []byte(password))

		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"message": "Contraseña incorrecta",
			})
			return
		}
		if clienteObteido.Activo != true {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"message": "no permitido",
			})
			return
		}

		key := []byte(os.Getenv("TOKEN"))

		token := jwt.New(jwt.SigningMethodHS256)

		claims := token.Claims.(jwt.MapClaims)

		claims["email"] = clienteObteido.Email

		claims["password"] = clienteObteido.Contrasena

		if clienteObteido.Email == "admin@admin.com" {
			claims["rol"] = "admin"
		} else {
			claims["rol"] = "cliente"
		}
		claims["exp"] = time.Now().Add(time.Hour * 3).Unix()

		tokenString, err := token.SignedString(key)

		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"message": "error al firmar el token: " + err.Error(),
			})
			return
		}

		ctx.JSON(http.StatusAccepted, gin.H{
			"data": tokenString,
		})

	}
}
