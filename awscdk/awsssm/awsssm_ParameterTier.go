package awsssm


// SSM parameter tier.
//
// Example:
//   ssm.NewStringParameter(this, jsii.String("Parameter"), &stringParameterProps{
//   	allowedPattern: jsii.String(".*"),
//   	description: jsii.String("The value Foo"),
//   	parameterName: jsii.String("FooParameter"),
//   	stringValue: jsii.String("Foo"),
//   	tier: ssm.parameterTier_ADVANCED,
//   })
//
type ParameterTier string

const (
	// String.
	ParameterTier_ADVANCED ParameterTier = "ADVANCED"
	// String.
	ParameterTier_INTELLIGENT_TIERING ParameterTier = "INTELLIGENT_TIERING"
	// String.
	ParameterTier_STANDARD ParameterTier = "STANDARD"
)

