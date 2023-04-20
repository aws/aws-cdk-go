package awsssm


// SSM parameter tier.
//
// Example:
//   ssm.NewStringParameter(this, jsii.String("Parameter"), &StringParameterProps{
//   	AllowedPattern: jsii.String(".*"),
//   	Description: jsii.String("The value Foo"),
//   	ParameterName: jsii.String("FooParameter"),
//   	StringValue: jsii.String("Foo"),
//   	Tier: ssm.ParameterTier_ADVANCED,
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

