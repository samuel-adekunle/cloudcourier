package cld

import (
	"fmt"

	file "github.com/Ibukun-tech/cloudcourier/File"
)

func (c *Cloudinary) UploadFile(fileInterface interface{}) error {
	files, ok := fileInterface.(file.File)
	if !ok {
		return fmt.Errorf("it must be of the type cloudcourier.File")
	}

	if err := files.CheckIfTheFileIsValid(); err != nil {
		return err
	}

	if !c.Connected() {
		return fmt.Errorf("no active cloudinary client")
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
