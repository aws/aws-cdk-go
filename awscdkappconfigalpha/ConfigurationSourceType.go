package awscdkappconfigalpha


// The configuration source type.
// Experimental.
type ConfigurationSourceType string

const (
	// Experimental.
	ConfigurationSourceType_S3 ConfigurationSourceType = "S3"
	// Experimental.
	ConfigurationSourceType_SECRETS_MANAGER ConfigurationSourceType = "SECRETS_MANAGER"
	// Experimental.
	ConfigurationSourceType_SSM_PARAMETER ConfigurationSourceType = "SSM_PARAMETER"
	// Experimental.
	ConfigurationSourceType_SSM_DOCUMENT ConfigurationSourceType = "SSM_DOCUMENT"
	// Experimental.
	ConfigurationSourceType_CODE_PIPELINE ConfigurationSourceType = "CODE_PIPELINE"
)

