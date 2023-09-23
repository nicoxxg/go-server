package turno

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nicoxxg/go-server/internal/domain/turno"
)

type TurnoController struct {
	turnoService turno.TurnoService
}

func NewTurnoController(service turno.TurnoService) *TurnoController {
	return &TurnoController{
		turnoService: service,
	}
}

func (c *TurnoController) Save() gin.HandlerFunc {

	return func(ctx *gin.Context) {
		var request turno.RequestTurno

		err := ctx.Bind(&request)

		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"message": "bad_request",
			})

			return
		}

		turnos, err := c.turnoService.Save(ctx, request)

		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"message": "internal_server_error",
			})

			return
		}

		ctx.JSON(http.StatusOK, gin.H{
			"data": turnos,
		})
	}

}

func (c *TurnoController) FindAllTurnos() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		turnos, err := c.turnoService.FindAll(ctx)
		fmt.Println(err)

		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"mensaje": "internal server error",
			})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{
			"data": turnos,
		})

	}
}
