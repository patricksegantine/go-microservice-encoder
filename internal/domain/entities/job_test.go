package entities_test

import (
	"encoder/internal/domain/entities"
	"testing"
	"time"

	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/require"
)

func TestNewJob(t *testing.T) {
	video := entities.NewVideo()
	video.ID = uuid.NewV4().String()
	video.FilePath = "path"
	video.CreatedAt = time.Now()

	job, err := entities.NewJob("path", "converted", video)
	require.NotNil(t, job)
	require.Nil(t, err)
}
