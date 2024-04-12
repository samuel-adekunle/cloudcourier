package config

const (
	CloudProviderCloudinary = "cloudinary"
	CloudProviderAWS        = "aws"
)

// Providers represent the cloud providers being supported at the moment.
//
// Expand this map as you add more providers.
var Providers = map[string]bool{
	CloudProviderCloudinary: true,
	CloudProviderAWS:        true,
}
