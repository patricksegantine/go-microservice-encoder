package repositories

import (
	"encoder/internal/domain/entities"
	"fmt"

	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
)

type VideoRepository struct {
	Db *gorm.DB
}

func NewVideoRepository(db *gorm.DB) *VideoRepository {
	return &VideoRepository{Db: db}
}

func (repo VideoRepository) Insert(video *entities.Video) (*entities.Video, error) {
	if video.ID == "" {
		video.ID = uuid.NewV4().String()
	}

	err := repo.Db.Create(video).Error
	if err != nil {
		return nil, err
	}

	return video, nil
}

func (repo VideoRepository) Find(id string) (*entities.Video, error) {
	var video entities.Video

	repo.Db.Preload("Jobs").First(&video, "id = ?", id)
	if video.ID == "" {
		return nil, fmt.Errorf("video does not exist")
	}

	return &video, nil
}
