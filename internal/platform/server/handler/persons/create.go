package persons

import (
	"net/http"

	xmen "xmen-mutant/internal"
	"xmen-mutant/internal/creating"
	"xmen-mutant/kit/command"

	"github.com/gin-gonic/gin"
)

type createRequest struct {
	Mutant bool     `json:"mutant"`
	Dna    []string `json:"dna"`
}

// CreateHandler returns an HTTP handler for persons creation.
func CreateHandler(commandBus command.Bus) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req createRequest
		if err := ctx.BindJSON(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, err.Error())
			return
		}

		err := commandBus.Dispatch(ctx, creating.NewPersonCommand(
			req.Mutant,
			req.Dna,
		))

		if err != nil {
			if err == xmen.ErrEmptyDna {
				ctx.JSON(http.StatusBadRequest, err.Error())
				return
			} else {
				ctx.JSON(http.StatusInternalServerError, err.Error())
				return
			}
		}

		ctx.Status(http.StatusCreated)
	}
}
