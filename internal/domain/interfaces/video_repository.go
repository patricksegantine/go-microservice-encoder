package interfaces

import (
	"encoder/internal/domain/entities"
)

type VideoRepository interface {
	Insert(video *entities.Video) (*entities.Video, error)
	Find(id string) (*entities.Video, error)
}
