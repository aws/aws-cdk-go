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
type ParameterTier string

const (
	// String.
	ParameterTier_ADVANCED ParameterTier = "ADVANCED"
	// String.
	ParameterTier_INTELLIGENT_TIERING ParameterTier = "INTELLIGENT_TIERING"
	// String.
	ParameterTier_STANDARD ParameterTier = "STANDARD"
)

