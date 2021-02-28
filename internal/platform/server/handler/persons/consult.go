package persons

import (
	"net/http"

	"xmen-mutant/internal/consulting"
	"xmen-mutant/kit/command"
	"xmen-mutant/kit/utils"

	"github.com/gin-gonic/gin"
)

// ConsultHandler returns an HTTP handler for persons search.
func ConsultHandler(commandBus command.Bus) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		result, err := commandBus.Dispatch(ctx, consulting.NewPersonCommand(
			0,
			false,
			nil,
		))
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, utils.CreateResponse(err))
			return
		}
		ctx.Status(http.StatusCreated)
		ctx.JSON(http.StatusOK, utils.CreateResponse(result))

	}
}
