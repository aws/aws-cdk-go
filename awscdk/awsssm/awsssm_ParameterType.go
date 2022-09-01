package awsssm


// SSM parameter type.
//
// Example:
//   // Example automatically generated from non-compiling source. May contain errors.
//   ssmParameter := ssm.NewStringParameter(this, jsii.String("mySsmParameter"), &stringParameterProps{
//   	parameterName: jsii.String("mySsmParameter"),
//   	stringValue: jsii.String("mySsmParameterValue"),
//   	type: ssm.parameterType_STRING,
//   })
//
//   secureParameter := ssm.NewStringParameter(this, jsii.String("mySecretParameter"), &stringParameterProps{
//   	parameterName: jsii.String("mySecretParameter"),
//   	stringValue: jsii.String("mySecretParameterValue"),
//   	type: ssm.*parameterType_SECURE_STRING,
//   })
//
//   listParameter := ssm.NewStringParameter(this, jsii.String("myListParameter"), &stringParameterProps{
//   	parameterName: jsii.String("myListParameter"),
//   	stringValue: []interface{}{
//   		jsii.String("myListParameterValue1"),
//   		jsii.String("myListParameterValue2"),
//   	},
//   	type: ssm.*parameterType_STRING_LIST,
//   })
//
type ParameterType string

const (
	// String.
	ParameterType_STRING ParameterType = "STRING"
	// Secure String.
	//
	// Parameter Store uses an AWS Key Management Service (KMS) customer master key (CMK) to encrypt the parameter value.
	// Parameters of type SecureString cannot be created directly from a CDK application.
	ParameterType_SECURE_STRING ParameterType = "SECURE_STRING"
	// String List.
	ParameterType_STRING_LIST ParameterType = "STRING_LIST"
	// An Amazon EC2 image ID, such as ami-0ff8a91507f77f867.
	ParameterType_AWS_EC2_IMAGE_ID ParameterType = "AWS_EC2_IMAGE_ID"
)

