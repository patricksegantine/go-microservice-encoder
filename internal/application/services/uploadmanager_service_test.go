package services_test

import (
	"encoder/internal/application/services"
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/require"
)

func init() {
	err := godotenv.Load("../../../configs/.env.local")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
}

func TestUploadManagerService_Upload(t *testing.T) {
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

	uploadmanager := services.NewUploadManagerService()
	uploadmanager.OutputBucket = OutputBucketName
	uploadmanager.VideoPath = os.Getenv("LOCAL_STORAGE_PATH") + "/" + video.ID

	doneUpload := make(chan string)
	go uploadmanager.ProcessUpload(50, doneUpload)

	result := <-doneUpload
	require.Equal(t, result, services.UploadCompleted)
}
