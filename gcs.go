package cloudcourier

import (
	"context"
	"fmt"
	"io"
	"path/filepath"

	"cloud.google.com/go/storage"
)

type GCSProviderConfig struct {
	Bucket string
}

func (c *GCSProviderConfig) GetProvider() Provider {
	return PROVIDER_GCS
}

type GCSProviderClient struct {
	Config *GCSProviderConfig
	Client *storage.Client
}

func newGCSProviderClient(config ProviderConfig) (ProviderClient, error) {
	gcsConfig, ok := config.(*GCSProviderConfig)
	if !ok {
		return nil, fmt.Errorf("incorrect provider config provided")
	}
	client, err := storage.NewClient(context.Background())
	if err != nil {
		return nil, err
	}
	return &GCSProviderClient{
		Config: gcsConfig,
		Client: client,
	}, nil
}

func init() {
	RegisterProviderConstructor(PROVIDER_GCS, newGCSProviderClient)
}

func (g *GCSProviderClient) UploadFile(filePath string, reader io.Reader) error {
	var BaseFileName string
	ctx := context.Background()
	if filePath != "" {
		BaseFileName = filepath.Base(filePath)
	}
	if filePath == "" {
		return fmt.Errorf("you did not specify the filepath")
	}
	obj := g.Client.Bucket(g.Config.Bucket).Object(BaseFileName)
	w := obj.NewWriter(ctx)
	if _, err := io.Copy(w, reader); err != nil {
		return fmt.Errorf("you did not set the reader to the file")
	}
	if err := w.Close(); err != nil {
		return fmt.Errorf("could not upload file")
	}
	v, err := obj.Attrs(ctx)
	if err != nil {
		return fmt.Errorf("%s", err)
	}
	fmt.Println(v)
	return nil
}

func (g *GCSProviderClient) DeleteFile(fieldID string) error {
	return fmt.Errorf("Unimplemented")
}
func (g *GCSProviderClient) ListFiles(directory string) ([]string, error) {
	return nil, fmt.Errorf("Unimplemented")
}

func (g *GCSProviderClient) GetFile(fileID string) (io.Reader, error) {
	return nil, fmt.Errorf("Unimplemented")

}
