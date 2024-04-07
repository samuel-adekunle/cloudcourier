package types

type CloudCourierBridge struct {
	// This helps to show the cloud provider you want to use
	CloudProvider string
	// The Api key for now I am just using it to access the cloudinary services
	ApiKey string
	// The Api secret which corresponds to the ApiKey
	ApiSecret string
	// This is for cloudinary you need to provide the cloud name for the cloudinary
	CloudName string
}
