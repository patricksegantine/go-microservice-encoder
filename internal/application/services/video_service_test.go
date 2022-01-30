package services_test

import (
	"encoder/internal/application/services"
	"encoder/internal/domain/entities"
	"encoder/internal/infrastructure/database"
	"encoder/internal/infrastructure/repositories"
	"fmt"
	"log"
	"os"
	"testing"
	"time"

	"github.com/joho/godotenv"
	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/require"
)

func init() {
	err := godotenv.Load("../../../configs/.env.local")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
}

func prepare() (*entities.Video, *repositories.VideoRepository) {
	db := database.NewDbTest()
	defer db.Close()

	video := entities.NewVideo()
	video.ID = uuid.NewV4().String()
	video.FilePath = "test/Fundamentos-Mensageria.mp4"
	video.CreatedAt = time.Now()

	repo := repositories.NewVideoRepository(db)
	return video, repo
}

func TestVideoService_DownloadAndFragmentAndEncode(t *testing.T) {
	video, repo := prepare()

	credentials := os.Getenv("GOOGLE_APPLICATION_CREDENTIALS")
	fmt.Println(credentials)

	service := services.NewVideoService()
	service.Video = video
	service.VideoRepository = repo

	err := service.Download("seganta-video-encoder")
	require.Nil(t, err)

	err = service.Fragment()
	require.Nil(t, err)

	err = service.Encode()
	require.Nil(t, err)

	err = service.Finish()
	require.Nil(t, err)
}
