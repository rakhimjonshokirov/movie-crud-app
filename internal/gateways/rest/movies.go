package rest

import (
	"rakhimjonshokirov/movie-crud-app/internal/entities"
	"rakhimjonshokirov/movie-crud-app/pkg/errs"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// @Summary Get all movies
// @Tags movies
// @Produce json
// @Router /movies [get]
// @Success 200 {object} R{data=entities.Movie}
// @Failure 401 {object} R
// @Failure 422 {object} R
// @Failure 500 {object} R
func (h *Handler) getAllMovies(c *gin.Context) {
	movies, err := h.movieUsc.GetAllMovies(c)
	if err != nil {
		h.l.Error("getAllMovies failed", zap.Error(err))

		Return(c, nil, err)
		return
	}

	Return(c, movies, nil)
}

// @Summary Get movie by ID
// @Tags movies
// @Produce json
// @Param id path int true "Movie ID"
// @Router /movies/{id} [get]
// @Success 200 {object} R{data=entities.Movie}
// @Failure 401 {object} R
// @Failure 422 {object} R
// @Failure 500 {object} R
func (h *Handler) getMovieByID(c *gin.Context) {
	idStr := c.Param("id")
	idUint, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		h.l.Error("strconv.ParseUint failed", zap.Error(err))

		Return(c, nil, errs.ErrValidation)
		return
	}

	movie, err := h.movieUsc.GetMovieByID(c, uint(idUint))
	if err != nil {
		h.l.Error("GetMovieByID failed", zap.Error(err))

		Return(c, nil, err)
		return
	}

	Return(c, movie, nil)
}

// @Summary Create a movie
// @Tags movies
// @Accept json
// @Produce json
// @Param movie body entities.Movie true "Movie data"
// @Router /movies [post]
// @Success 200 {object} R{data=entities.Movie}
// @Failure 401 {object} R
// @Failure 422 {object} R
// @Failure 500 {object} R
func (h *Handler) createMovie(c *gin.Context) {
	var movie entities.Movie
	if err := c.ShouldBindJSON(&movie); err != nil {
		Return(c, nil, err)
		return
	}
	if err := h.movieUsc.CreateMovie(c, &movie); err != nil {
		h.l.Error("GetMovieByID failed", zap.Error(err))

		Return(c, nil, err)
		return
	}

	Return(c, nil, nil)
}

// @Summary Update movie
// @Tags movies
// @Accept json
// @Produce json
// @Param id path int true "Movie ID"
// @Param movie body entities.Movie true "Movie data"
// @Router /movies/{id} [put]
// @Success 200 {object} R{data=entities.Movie}
// @Failure 401 {object} R
// @Failure 422 {object} R
// @Failure 500 {object} R
func (h *Handler) updateMovie(c *gin.Context) {
	var movie entities.Movie
	id, _ := strconv.Atoi(c.Param("id"))
	if err := c.ShouldBindJSON(&movie); err != nil {

		Return(c, nil, err)
		return
	}
	movie.ID = uint(id)
	if err := h.movieUsc.UpdateMovie(c, &movie); err != nil {
		h.l.Error("UpdateMovie failed", zap.Error(err))

		Return(c, nil, err)
		return
	}

	Return(c, nil, nil)
}

// @Summary Delete movie
// @Tags movies
// @Produce json
// @Param id path int true "Movie ID"
// @Router /movies/{id} [delete]
// @Success 200 {object} R
// @Failure 401 {object} R
// @Failure 422 {object} R
// @Failure 500 {object} R
func (h *Handler) deleteMovie(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	err := h.movieUsc.DeleteMovie(c, uint(id))
	if err != nil {
		h.l.Error("DeleteMovie failed", zap.Error(err))

		Return(c, nil, nil)
		return
	}

	Return(c, nil, nil)
}
