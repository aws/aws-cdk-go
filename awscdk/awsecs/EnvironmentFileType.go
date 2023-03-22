package awsecs


// Type of environment file to be included in the container definition.
type EnvironmentFileType string

const (
	// Environment file hosted on S3, referenced by object ARN.
	EnvironmentFileType_S3 EnvironmentFileType = "S3"
)

