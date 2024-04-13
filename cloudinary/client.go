package cloudinary_client

import "io"

// TODO: model this properly, implement the functions properly too.

type CloudinaryClient struct {
}

func (client CloudinaryClient) UploadFile(filePath string, reader io.Reader) error {
	//TODO implement me
	panic("implement me")
}

func (client CloudinaryClient) DeleteFile(fileID string) error {
	//TODO implement me
	panic("implement me")
}

func (client CloudinaryClient) ListFiles(directory string) ([]string, error) {
	//TODO implement me
	panic("implement me")
}

func (client CloudinaryClient) GetFile(fileID string) (io.Reader, error) {
	//TODO implement me
	panic("implement me")
}
