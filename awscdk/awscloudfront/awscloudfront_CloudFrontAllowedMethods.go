package awscloudfront


// An enum for the supported methods to a CloudFront distribution.
// Experimental.
type CloudFrontAllowedMethods string

const (
	// Experimental.
	CloudFrontAllowedMethods_GET_HEAD CloudFrontAllowedMethods = "GET_HEAD"
	// Experimental.
	CloudFrontAllowedMethods_GET_HEAD_OPTIONS CloudFrontAllowedMethods = "GET_HEAD_OPTIONS"
	// Experimental.
	CloudFrontAllowedMethods_ALL CloudFrontAllowedMethods = "ALL"
)

