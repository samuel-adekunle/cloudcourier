package cld

import (
	"fmt"
	"os"
	"path/filepath"

	file "github.com/Ibukun-tech/cloudcourier/File"
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
	var err error
	if files.Path != "" {
		_, err = os.Stat(files.Path)
		if err != nil {
			return fmt.Errorf("this path does not exist")
		}
		if !os.IsNotExist(err) {
			return fmt.Errorf("this path does not exist")
		}
		file, err := os.Open()
		if err != nil {
			return fmt.Errorf("this file does not exist")
		}
		defer file.Close()
		files.ToHandle = file

		if files.FileName == "" {
			files.FileName = filepath.Base(files.Path)
		}
	}
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
