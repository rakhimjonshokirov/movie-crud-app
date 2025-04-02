package rest

import (
	"rakhimjonshokirov/movie-crud-app/pkg/errs"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type registerReq struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

// @Summary      Register a new user
// @Description  Creates a new user with email, username and password
// @Router       /auth/register [post]
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        request body registerReq true "req body"
// @Success      200 {object} R{data=string}
// @Failure      400 {object} R
// @Failure      422 {object} R
// @Failure      500 {object} R
func (h *Handler) register(c *gin.Context) {
	req := registerReq{}
	if err := c.ShouldBindJSON(&req); err != nil {
		h.l.Error("register validation failed", zap.Error(err))

		Return(c, nil, errs.ErrValidation)
		return
	}

	if err := h.usersUsc.Register(req.Email, req.Username, req.Password); err != nil {
		h.l.Error("h.usersUsc.Register failed", zap.Error(err))

		Return(c, nil, err)
		return
	}

	Return(c, "registered successfully", nil)
}

type loginReq struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// @Summary      Login
// @Description  Logs in a user and returns a JWT token
// @Router       /auth/login [post]
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        request body loginReq true "req body"
// @Success      200 {object} R{data=RespLogin}
// @Failure      400 {object} R
// @Failure      401 {object} R
// @Failure      422 {object} R
// @Failure      500 {object} R
func (h *Handler) login(c *gin.Context) {
	req := loginReq{}

	if err := c.ShouldBindJSON(&req); err != nil {
		h.l.Error("login validation failed", zap.Error(err))

		Return(c, nil, errs.ErrValidation)
		return
	}

	token, err := h.usersUsc.Login(req.Email, req.Password)
	if err != nil {
		h.l.Error("h.usersUsc.Login failed", zap.Error(err))

		Return(c, nil, err)
		return
	}

	Return(c, RespLogin{token}, nil)
}

type RespLogin struct {
	Token string `json:"token"`
}
