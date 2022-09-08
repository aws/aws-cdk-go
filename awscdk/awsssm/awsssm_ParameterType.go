package awsssm


// SSM parameter type.
// Experimental.
type ParameterType string

const (
	// String.
	// Experimental.
	ParameterType_STRING ParameterType = "STRING"
	// Secure String.
	//
	// Parameter Store uses an AWS Key Management Service (KMS) customer master key (CMK) to encrypt the parameter value.
	// Parameters of type SecureString cannot be created directly from a CDK application.
	// Experimental.
	ParameterType_SECURE_STRING ParameterType = "SECURE_STRING"
	// String List.
	// Experimental.
	ParameterType_STRING_LIST ParameterType = "STRING_LIST"
	// An Amazon EC2 image ID, such as ami-0ff8a91507f77f867.
	// Experimental.
	ParameterType_AWS_EC2_IMAGE_ID ParameterType = "AWS_EC2_IMAGE_ID"
)

