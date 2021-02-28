package persons

import (
	"fmt"
	"net/http"

	"xmen-mutant/internal/consulting"
	"xmen-mutant/kit/command"

	"github.com/gin-gonic/gin"
)

// ConsultHandler returns an HTTP handler for persons search.
func ConsultHandler(commandBus command.Bus) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		result, err := commandBus.Dispatch(ctx, consulting.NewPersonCommand(
			3,
			false,
			[]string{"ATGCGA", "CAGTGC", "TTATGT", "AGAAGG", "CCCCTA", "TCACTG"},
		))
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, err.Error())
			return
		}
		fmt.Println(result)
		ctx.Status(http.StatusCreated)
	}
}
