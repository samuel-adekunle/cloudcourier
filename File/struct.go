package file

import (
	"errors"
	"io"
)

type File struct {
	// To handle the file
	ToHandle io.Reader
	// This is the path of the file
	Path string `json:"path"`
	// This is the file name of the file to be stored
	FileName string `json:"fileName"`
}

type FormFile struct {
	Files []*File `json:"multiFiles"`
}

func (f *File) CheckIfTheFileIsValid() error {
	if f.Path == "" && f.ToHandle == nil {
		return errors.New("you have to set either the path or the handle")
	}
	if f.Path != "" && f.ToHandle != nil {
		return errors.New("you can only set once at a time either the path or tohandle")
	}
	if f.FileName == "" && f.ToHandle != nil {
		return errors.New("before you set the handle you need to specify the name of the file")
	}
	return nil
}
