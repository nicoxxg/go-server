package cliente

import (
	"net/http"
	"strconv"

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

func (c *ClienteController) UpdateClient() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var request cliente.ClientRequest

		err := ctx.Bind(&request)

		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"message": "bad_request",
			})
			return
		}

		clienteObtenido, err := c.clienteService.UpdateClient(ctx, request)

		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"message": "internal_server_error",
				"error":   err.Error(),
			})

			return
		}
		ctx.JSON(http.StatusOK, gin.H{
			"data": clienteObtenido,
		})

	}

}

func (c *ClienteController) SaveClient() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var request cliente.ClientRequest

		err := ctx.Bind(&request)

		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"message": "bad_request",
			})

			return
		}

		clienteObtenido, err := c.clienteService.SaveClient(ctx, request)

		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"message": "internal_server_error",
				"error":   err.Error(),
			})

			return
		}
		ctx.JSON(http.StatusOK, gin.H{
			"data": clienteObtenido,
		})

	}
}

func (c *ClienteController) FindClienteByEmail() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		email := ctx.Query("email")

		if email == "" {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"mensaje": "emial vacio",
			})
			return
		}
		cliente, err := c.clienteService.FindByEmail(ctx, email)

		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"mensaje": "internal_server_error",
				"error":   err.Error(),
			})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{
			"data": cliente,
		})
	}
}

func (c *ClienteController) FindClientById() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		idParam, err := strconv.Atoi(ctx.Param("id"))

		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"mensaje": "id invalido",
			})
			return
		}
		id := int64(idParam)

		clienteObtenido, err := c.clienteService.FindById(ctx, id)

		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"mensaje": "internal_server_error",
				"error":   err.Error(),
			})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{
			"data": clienteObtenido,
		})

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
