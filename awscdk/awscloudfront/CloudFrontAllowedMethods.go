package awscloudfront


// An enum for the supported methods to a CloudFront distribution.
type CloudFrontAllowedMethods string

const (
	CloudFrontAllowedMethods_GET_HEAD CloudFrontAllowedMethods = "GET_HEAD"
	CloudFrontAllowedMethods_GET_HEAD_OPTIONS CloudFrontAllowedMethods = "GET_HEAD_OPTIONS"
	CloudFrontAllowedMethods_ALL CloudFrontAllowedMethods = "ALL"
)

