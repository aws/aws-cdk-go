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
// Experimental.
type ParameterTier string

const (
	// String.
	// Experimental.
	ParameterTier_ADVANCED ParameterTier = "ADVANCED"
	// String.
	// Experimental.
	ParameterTier_INTELLIGENT_TIERING ParameterTier = "INTELLIGENT_TIERING"
	// String.
	// Experimental.
	ParameterTier_STANDARD ParameterTier = "STANDARD"
)

