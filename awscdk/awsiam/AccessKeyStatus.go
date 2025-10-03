package awsiam


// Valid statuses for an IAM Access Key.
type AccessKeyStatus string

const (
	// An active access key.
	//
	// An active key can be used to make API calls.
	AccessKeyStatus_ACTIVE AccessKeyStatus = "ACTIVE"
	// An inactive access key.
	//
	// An inactive key cannot be used to make API calls.
	AccessKeyStatus_INACTIVE AccessKeyStatus = "INACTIVE"
)

