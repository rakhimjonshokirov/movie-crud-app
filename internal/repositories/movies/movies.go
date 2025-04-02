package movies

import (
	"rakhimjonshokirov/movie-crud-app/internal/entities"

	"gorm.io/gorm"
)

type MovieRepository struct {
	db *gorm.DB
}

func NewMovieRepository(db *gorm.DB) *MovieRepository {
	return &MovieRepository{db}
}

func (r *MovieRepository) Create(movie *entities.Movie) error {
	return r.db.Create(movie).Error
}

func (r *MovieRepository) FindAll() ([]entities.Movie, error) {
	var movies []entities.Movie
	err := r.db.Find(&movies).Error
	return movies, err
}

func (r *MovieRepository) FindByID(id uint) (*entities.Movie, error) {
	var movie entities.Movie

	tx := r.db.First(&movie, id)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return &movie, nil
}

func (r *MovieRepository) Update(movie *entities.Movie) error {
	return r.db.Save(movie).Error
}

func (r *MovieRepository) Delete(movie *entities.Movie) error {
	return r.db.Delete(movie).Error
}
