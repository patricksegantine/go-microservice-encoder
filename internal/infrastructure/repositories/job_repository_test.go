package repositories_test

import (
	"encoder/internal/domain/entities"
	"encoder/internal/infrastructure/database"
	"encoder/internal/infrastructure/repositories"
	"testing"
	"time"

	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/require"
)

var (
	db        *gorm.DB
	repoVideo *repositories.VideoRepository
	repoJob   *repositories.JobRepository
)

func prepare() {
	db = database.NewDbTest()
	repoVideo = repositories.NewVideoRepository(db)
	repoJob = repositories.NewJobRepository(db)
}

func TestJobRepositoryInsertAndFind(t *testing.T) {
	prepare()
	defer db.Close()

	video := createVideo()
	repoVideo.Insert(video)

	job, err := entities.NewJob("output_path", "Pending", video)
	require.Nil(t, err)

	repoJob.Insert(job)

	expectedJob, err := repoJob.Find(job.ID)
	require.NotEmpty(t, expectedJob.ID)
	require.Nil(t, err)
	require.Equal(t, expectedJob.ID, job.ID)
	require.Equal(t, expectedJob.VideoID, video.ID)
}

func TestJobRepositoryUpdate(t *testing.T) {
	prepare()
	defer db.Close()

	video := createVideo()
	repoVideo.Insert(video)

	job, err := entities.NewJob("output_path", "Pending", video)
	require.Nil(t, err)

	repoJob.Insert(job)
	job.Status = entities.JobStatusComplete
	repoJob.Update(job)

	result, err := repoJob.Find(job.ID)
	require.NotEmpty(t, result.ID)
	require.Nil(t, err)
	require.Equal(t, result.Status, job.Status)
}

func createVideo() *entities.Video {
	video := entities.NewVideo()
	video.ID = uuid.NewV4().String()
	video.FilePath = "path"
	video.CreatedAt = time.Now()

	return video
}
