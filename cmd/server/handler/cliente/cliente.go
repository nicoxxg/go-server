package cliente

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nicoxxg/go-server/internal/domain/cliente"
)

type ClienteController struct {
	clienteService cliente.ClientService
}

func NewClientController(service cliente.ClientService) *ClienteController {
	return &ClienteController{
		clienteService: service,
	}
}

func (c *ClienteController) FindAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		clientes, err := c.clienteService.FindAll(ctx)

		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"mensaje": "internal server error",
			})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{
			"data": clientes,
		})

	}
}
