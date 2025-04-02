package rest

import (
	"github.com/gin-gonic/gin"
)

func (s *Handler) ping(c *gin.Context) {
	Return(c, "ping", nil)
}
