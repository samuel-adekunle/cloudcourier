package types

type CloudCourierBridge struct {
	// This helps to show the cloud provider you want to use
	Provider string
	// The Api key for now I am just using it to access the cloudinary services
	ApiKey string
	// The Api secret which corresponds to the ApiKey
	ApiSecret string
}
