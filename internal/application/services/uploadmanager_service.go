package services

import (
	"context"
	"io"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"strings"

	"cloud.google.com/go/storage"
)

const (
	UploadCompleted = "upload completed"
)

type UploadManagerService struct {
	VideoPath    string
	Files        []string
	OutputBucket string
	Errors       []string
}

func NewUploadManagerService() *UploadManagerService {
	return &UploadManagerService{}
}

// UploadObject sends a object to Google Cloud bucket
func (um *UploadManagerService) UploadObject(objectPath string, client *storage.Client, ctx context.Context) error {
	path := strings.Split(objectPath, os.Getenv("LOCAL_STORAGE_PATH")+"/")

	f, err := os.Open(objectPath)
	if err != nil {
		return err
	}

	defer f.Close()

	wc := client.Bucket((um.OutputBucket)).Object(path[1]).NewWriter(ctx)
	// wc.ACL = []storage.ACLRule{
	// 	{
	// 		Entity: storage.AllUsers,
	// 		Role:   storage.RoleReader,
	// 	},
	// }

	if _, err = io.Copy(wc, f); err != nil {
		return err
	}

	if err := wc.Close(); err != nil {
		return err
	}

	return nil
}

// ProcessUpload starts the process of file upload
func (um *UploadManagerService) ProcessUpload(concurrency int, doneUpload chan string) error {
	in := make(chan int, runtime.NumCPU())
	returnChannel := make(chan string)

	err := um.loadFilePaths()
	if err != nil {
		return err
	}

	uploadClient, ctx, err := getClientUpload()
	if err != nil {
		return err
	}

	for process := 0; process < concurrency; process++ {
		go um.uploadWorker(in, returnChannel, uploadClient, ctx)
	}

	go func() {
		for x := 0; x < len(um.Files); x++ {
			in <- x
		}
		close(in)
	}()

	for r := range returnChannel {
		if r != "" {
			doneUpload <- r
			break
		}
	}

	return nil
}

func (um *UploadManagerService) loadFilePaths() error {
	err := filepath.Walk(um.VideoPath, func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			um.Files = append(um.Files, path)
		}
		return nil
	})

	if err != nil {
		return err
	}
	return nil
}

func (um *UploadManagerService) uploadWorker(in chan int, returnChan chan string, uploadClient *storage.Client, ctx context.Context) {
	for x := range in {
		err := um.UploadObject(um.Files[x], uploadClient, ctx)

		if err != nil {
			um.Errors = append(um.Errors, um.Files[x])
			log.Printf("error during the upload file: %v. Error: %v", um.Files[x], err)
			returnChan <- err.Error()
		}

		returnChan <- ""
	}

	returnChan <- UploadCompleted
}

func getClientUpload() (*storage.Client, context.Context, error) {
	ctx := context.Background()

	client, err := storage.NewClient(ctx)
	if err != nil {
		return nil, nil, err
	}
	return client, ctx, nil
}
