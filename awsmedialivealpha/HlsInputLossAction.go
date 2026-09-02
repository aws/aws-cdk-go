package awsmedialivealpha

import (
	_init_ "github.com/aws/aws-cdk-go/awsmedialivealpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// HLS input loss action.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import medialive_alpha "github.com/aws/aws-cdk-go/awsmedialivealpha"
//
//   hlsInputLossAction := medialive_alpha.HlsInputLossAction_Of(jsii.String("value"))
//
// Experimental.
type HlsInputLossAction interface {
	// The underlying string value passed to CloudFormation.
	// Experimental.
	Value() *string
}

// The jsii proxy struct for HlsInputLossAction
type jsiiProxy_HlsInputLossAction struct {
	_ byte // padding
}

func (j *jsiiProxy_HlsInputLossAction) Value() *string {
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
func HlsInputLossAction_Of(value *string) HlsInputLossAction {
	_init_.Initialize()

	if err := validateHlsInputLossAction_OfParameters(value); err != nil {
		panic(err)
	}
	var returns HlsInputLossAction

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-medialive-alpha.HlsInputLossAction",
		"of",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func HlsInputLossAction_EMIT_OUTPUT() HlsInputLossAction {
	_init_.Initialize()
	var returns HlsInputLossAction
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.HlsInputLossAction",
		"EMIT_OUTPUT",
		&returns,
	)
	return returns
}

func HlsInputLossAction_PAUSE_OUTPUT() HlsInputLossAction {
	_init_.Initialize()
	var returns HlsInputLossAction
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.HlsInputLossAction",
		"PAUSE_OUTPUT",
		&returns,
	)
	return returns
}

