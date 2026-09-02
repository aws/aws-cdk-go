package awsmedialivealpha

import (
	_init_ "github.com/aws/aws-cdk-go/awsmedialivealpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Feature activation state.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import medialive_alpha "github.com/aws/aws-cdk-go/awsmedialivealpha"
//
//   featureActivationState := medialive_alpha.FeatureActivationState_Of(jsii.String("value"))
//
// Experimental.
type FeatureActivationState interface {
	// The underlying string value passed to CloudFormation.
	// Experimental.
	Value() *string
}

// The jsii proxy struct for FeatureActivationState
type jsiiProxy_FeatureActivationState struct {
	_ byte // padding
}

func (j *jsiiProxy_FeatureActivationState) Value() *string {
	var returns *string
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}


// A value not yet modelled by AWS CDK.
// Experimental.
func FeatureActivationState_Of(value *string) FeatureActivationState {
	_init_.Initialize()

	if err := validateFeatureActivationState_OfParameters(value); err != nil {
		panic(err)
	}
	var returns FeatureActivationState

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-medialive-alpha.FeatureActivationState",
		"of",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func FeatureActivationState_DISABLED() FeatureActivationState {
	_init_.Initialize()
	var returns FeatureActivationState
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.FeatureActivationState",
		"DISABLED",
		&returns,
	)
	return returns
}

func FeatureActivationState_ENABLED() FeatureActivationState {
	_init_.Initialize()
	var returns FeatureActivationState
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.FeatureActivationState",
		"ENABLED",
		&returns,
	)
	return returns
}

