package awsssm


// SSM parameter data type.
type ParameterDataType string

const (
	// Text.
	ParameterDataType_TEXT ParameterDataType = "TEXT"
	// Aws Ec2 Image.
	ParameterDataType_AWS_EC2_IMAGE ParameterDataType = "AWS_EC2_IMAGE"
)

