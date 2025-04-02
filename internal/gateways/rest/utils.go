package rest

import (
	"errors"
	"net/http"

	"rakhimjonshokirov/movie-crud-app/pkg/status"

	"rakhimjonshokirov/movie-crud-app/pkg/errs"

	"github.com/gin-gonic/gin"
)

func (s *Handler) respond(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, R{
		Status:    status.Success,
		ErrorCode: 0,
		ErrorNote: "",
		Data:      data,
	})
}

func Return(c *gin.Context, data interface{}, err error) {
	switch {
	case err == nil:
		c.JSON(http.StatusOK, R{
			Status:    status.Success,
			ErrorCode: status.NoError,
			ErrorNote: "",
			Data:      data,
		})
	case errors.Is(err, errs.ErrValidation):
		c.JSON(http.StatusUnprocessableEntity, R{
			Status:    status.Failure,
			ErrorCode: status.ErrorCodeValidation,
			ErrorNote: err.Error(),
		})
	default:
		c.JSON(http.StatusInternalServerError, R{
			Status:    status.Failure,
			ErrorCode: status.ErrorCodeUniversal,
			ErrorNote: err.Error(),
		})
	}
}
