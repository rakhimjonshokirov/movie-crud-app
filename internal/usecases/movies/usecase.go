package movies

import (
	"context"
	"rakhimjonshokirov/movie-crud-app/internal/entities"

	"go.uber.org/zap"
)

type movieRepo interface {
	FindAll() ([]entities.Movie, error)
	FindByID(id uint) (*entities.Movie, error)
	Update(movie *entities.Movie) error
	Delete(movie *entities.Movie) error
	Create(movie *entities.Movie) error
}

type UseCase struct {
	movieRepo movieRepo
	log       *zap.Logger
}

func newUseCase(
	repo movieRepo,
	log *zap.Logger,
) *UseCase {
	return &UseCase{
		movieRepo: repo,
		log:       log,
	}
}

func (uc *UseCase) GetMovieByID(ctx context.Context, id uint) (*entities.Movie, error) {
	uc.log.Info("GetMovieByID called", zap.Uint("id", id))
	movie, err := uc.movieRepo.FindByID(id)
	if err != nil {
		uc.log.Error("failed to find movie by ID", zap.Error(err), zap.Uint("id", id))
		return nil, err
	}
	return movie, nil
}

func (uc *UseCase) GetAllMovies(ctx context.Context) ([]entities.Movie, error) {
	uc.log.Info("GetAllMovies called")
	movies, err := uc.movieRepo.FindAll()
	if err != nil {
		uc.log.Error("failed to fetch movies", zap.Error(err))
		return nil, err
	}
	return movies, nil
}

func (uc *UseCase) CreateMovie(ctx context.Context, movie *entities.Movie) error {
	uc.log.Info("CreateMovie called", zap.String("title", movie.Title))
	err := uc.movieRepo.Create(movie)
	if err != nil {
		uc.log.Error("failed to create movie", zap.Error(err))
	}
	return err
}

func (uc *UseCase) UpdateMovie(ctx context.Context, movie *entities.Movie) error {
	uc.log.Info("UpdateMovie called", zap.Uint("id", movie.ID))
	err := uc.movieRepo.Update(movie)
	if err != nil {
		uc.log.Error("failed to update movie", zap.Error(err), zap.Uint("id", movie.ID))
	}
	return err
}

func (uc *UseCase) DeleteMovie(ctx context.Context, id uint) error {
	uc.log.Info("DeleteMovie called", zap.Uint("id", id))
	movie, err := uc.movieRepo.FindByID(id)
	if err != nil {
		uc.log.Error("failed to find movie before delete", zap.Error(err))
		return err
	}
	err = uc.movieRepo.Delete(movie)
	if err != nil {
		uc.log.Error("failed to delete movie", zap.Error(err))
	}
	return err
}
