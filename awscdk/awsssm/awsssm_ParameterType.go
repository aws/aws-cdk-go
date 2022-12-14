package awsssm


// SSM parameter type.
// Deprecated: these types are no longer used.
type ParameterType string

const (
	// String.
	// Deprecated: these types are no longer used.
	ParameterType_STRING ParameterType = "STRING"
	// Secure String.
	//
	// Parameter Store uses an AWS Key Management Service (KMS) customer master key (CMK) to encrypt the parameter value.
	// Parameters of type SecureString cannot be created directly from a CDK application.
	// Deprecated: these types are no longer used.
	ParameterType_SECURE_STRING ParameterType = "SECURE_STRING"
	// String List.
	// Deprecated: these types are no longer used.
	ParameterType_STRING_LIST ParameterType = "STRING_LIST"
	// An Amazon EC2 image ID, such as ami-0ff8a91507f77f867.
	// Deprecated: these types are no longer used.
	ParameterType_AWS_EC2_IMAGE_ID ParameterType = "AWS_EC2_IMAGE_ID"
)

