package repositories

import (
	"encoder/internal/domain/entities"
	"fmt"

	"github.com/jinzhu/gorm"
)

type JobRepository struct {
	Db *gorm.DB
}

func NewJobRepository(db *gorm.DB) *JobRepository {
	return &JobRepository{Db: db}
}

func (repo *JobRepository) Insert(job *entities.Job) (*entities.Job, error) {
	err := repo.Db.Create(job).Error

	if err != nil {
		return nil, err
	}

	return job, nil
}

func (repo *JobRepository) Find(id string) (*entities.Job, error) {
	var job entities.Job

	repo.Db.Preload("Video").First(&job, "id = ?", id)
	if job.ID == "" {
		return nil, fmt.Errorf("job does not exist")
	}

	return &job, nil
}

func (repo *JobRepository) Update(job *entities.Job) (*entities.Job, error) {
	err := repo.Db.Save(&job).Error

	if err != nil {
		return nil, err
	}

	return job, nil
}
