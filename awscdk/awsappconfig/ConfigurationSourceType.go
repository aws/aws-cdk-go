package awsappconfig


// The configuration source type.
type ConfigurationSourceType string

const (
	ConfigurationSourceType_S3 ConfigurationSourceType = "S3"
	ConfigurationSourceType_SECRETS_MANAGER ConfigurationSourceType = "SECRETS_MANAGER"
	ConfigurationSourceType_SSM_PARAMETER ConfigurationSourceType = "SSM_PARAMETER"
	ConfigurationSourceType_SSM_DOCUMENT ConfigurationSourceType = "SSM_DOCUMENT"
	ConfigurationSourceType_CODE_PIPELINE ConfigurationSourceType = "CODE_PIPELINE"
)

