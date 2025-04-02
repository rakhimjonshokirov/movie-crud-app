package rest

import (
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func (h *Handler) routes() {
	url := ginSwagger.URL("/swagger/doc.json")
	h.router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	// ping
	h.router.GET("ping", h.ping)

	// Public routes
	authU := h.router.Group("/auth")
	authU.POST("/register", h.register)
	authU.POST("/login", h.login)

	authGroup := h.router.Group("/movies")
	authGroup.Use(JWTMiddleware()) // JWT middleware
	{
		authGroup.POST("", h.createMovie)
		authGroup.PUT("/:id", h.updateMovie)
		authGroup.DELETE("/:id", h.deleteMovie)
	}

	// Public routes
	h.router.GET("/movies", h.getAllMovies)
	h.router.GET("/movies/:id", h.getMovieByID)

}
