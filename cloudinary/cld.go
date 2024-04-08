package cld

import (
	"context"
	"fmt"
	"os"
	"path/filepath"

	file "github.com/Ibukun-tech/cloudcourier/File"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
)

func (c *Cloudinary) UploadFile(fileInterface interface{}) error {
	files, ok := fileInterface.(file.File)
	if !ok {
		return fmt.Errorf("it must be of the type cloudcourier.File")
	}
	// To check if its if the file fits the requirement
	if err := files.CheckIfTheFileIsValid(); err != nil {
		return err
	}

	if !c.Connected() {
		return fmt.Errorf("no active cloudinary client")
	}
	// var err error
	if files.Path != "" {
		_, err = os.Stat(files.Path)
		if err != nil {
			return fmt.Errorf("this path does not exist")
		}
		if !os.IsNotExist(err) {
			return fmt.Errorf("this path does not exist")
		}
		file, err := os.Open(files.Path)
		if err != nil {
			return fmt.Errorf("this file does not exist")
		}
		defer file.Close()
		files.ToHandle = file

		if files.FileName == "" {
			files.FileName = filepath.Base(files.Path)
		}
	}
	var ctx context.Context;
	// Work on getting a random public Id or I can specify it from the
	// I feel i need to implement the side like a go routine why because I am upload the file into the storage activite and also trying to get resources from it
	resp err := c.Client.Upload.Upload(ctx, files.FileName, uploader.UploadParams{PublicID: c.CloudName})
	if err != nil {
		return fmt.Errorf("failed to upload to the cloud storage %s", err)
	}

	fmt.Println(resp.URL)
	return nil
}

func (c *Cloudinary) Connected() bool {
	return c.Client != nil
}

func (c *Cloudinary) DisConnect() error {
	if c.Client != nil {
		c.Client = nil
	}
	return nil
}
