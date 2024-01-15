package controllers

import (
	"context"
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"time"
	"strconv"
	"cloud.google.com/go/storage"
	"path/filepath"
)

const (
	projectID  = "pariwarain"  // FILL IN WITH YOURS
	bucketName = "pariwarain" // FILL IN WITH YOURS
)

type ClientUploader struct {
	cl         *storage.Client
	projectID  string
	bucketName string
	uploadPath string
}

var uploader *ClientUploader = setUploader()

func setUploader() *ClientUploader {
	os.Setenv("GOOGLE_APPLICATION_CREDENTIALS", "./pariwarain-1bf55ccd268e.json") // FILL IN WITH YOUR FILE PATH
	client, err := storage.NewClient(context.Background())
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	upload := &ClientUploader{
		cl:         client,
		bucketName: bucketName,
		projectID:  projectID,
		uploadPath: "user/",
	}

	return upload 
}

func UploadImage (w http.ResponseWriter, r *http.Request) string {

	err := r.ParseMultipartForm(10 << 20) // 10 MB limit for file size
	if err != nil {
		http.Error(w, "Unable to parse form", http.StatusBadRequest)
		// panic(err)
		return ""
	}

	// Get a reference to the file from the request
	file, handler, err := r.FormFile("file")
	if err != nil {
		http.Error(w, "Error retrieving the file", http.StatusBadRequest)
		return ""
	}
	defer file.Close()

	if err != nil {
		// panic(err)
		return ""
	}

	currentTime := time.Now()
	
	filePath := handler.Filename
	fileExtension := filepath.Ext(filePath)

	fileName := filepath.Base(filePath)
	fileNameWithoutExt := fileName[:len(fileName)-len(filepath.Ext(fileName))]

	var name string = fileNameWithoutExt + strconv.FormatInt(currentTime.UnixNano(), 10) + fileExtension

	err = uploader.UploadFile(file, name)
	if err != nil {
		panic(err)
	}

	var url string = "https://storage.googleapis.com/pariwarain/user/" + name

	return url
}

func (c *ClientUploader) UploadFile(file multipart.File, object string) error {
	ctx := context.Background()

	ctx, cancel := context.WithTimeout(ctx, time.Second*50)
	defer cancel()

	// Upload an object with storage.Writer.
	wc := c.cl.Bucket(c.bucketName).Object(c.uploadPath + object).NewWriter(ctx)
	if _, err := io.Copy(wc, file); err != nil {
		return fmt.Errorf("io.Copy: %v", err)
	}
	if err := wc.Close(); err != nil {
		return fmt.Errorf("Writer.Close: %v", err)
	}

	return nil
}