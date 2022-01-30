package entities_test

import (
	"encoder/internal/domain/entities"
	"testing"
	"time"

	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/require"
)

func TestValidateIfVideoIsEmpty(t *testing.T) {
	video := entities.NewVideo()
	err := video.Validate()

	require.Error(t, err)
}

func TestVideoWhenIdIsUuid(t *testing.T) {
	video := entities.NewVideo()
	video.ID = uuid.NewV4().String()
	video.ResourceID = "Video resource id"
	video.FilePath = "path"
	video.CreatedAt = time.Now()

	err := video.Validate()

	require.Nil(t, err)
}

func TestVideoWhenIdIsNotUuid(t *testing.T) {
	video := entities.NewVideo()
	video.ID = "123456"
	video.ResourceID = "Video resource id"
	video.FilePath = "path"
	video.CreatedAt = time.Now()

	err := video.Validate()

	require.Error(t, err)
}
