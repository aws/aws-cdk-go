package awsssm


// SSM parameter data type.
// Experimental.
type ParameterDataType string

const (
	// Text.
	// Experimental.
	ParameterDataType_TEXT ParameterDataType = "TEXT"
	// Aws Ec2 Image.
	// Experimental.
	ParameterDataType_AWS_EC2_IMAGE ParameterDataType = "AWS_EC2_IMAGE"
)

