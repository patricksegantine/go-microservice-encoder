package repositories_test

import (
	"encoder/internal/domain/entities"
	"encoder/internal/infrastructure/database"
	"encoder/internal/infrastructure/repositories"
	"testing"
	"time"

	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/require"
)

func TestVideoRepositoryInsertAndFind(t *testing.T) {
	db := database.NewDbTest()
	defer db.Close()

	video := entities.NewVideo()
	video.ID = uuid.NewV4().String()
	video.FilePath = "path"
	video.CreatedAt = time.Now()

	repo := repositories.NewVideoRepository(db)
	repo.Insert(video)

	v, err := repo.Find(video.ID)

	require.NotEmpty(t, v.ID)
	require.Nil(t, err)
	require.Equal(t, v.ID, video.ID)
}
