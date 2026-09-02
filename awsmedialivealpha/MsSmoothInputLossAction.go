package awsmedialivealpha

import (
	_init_ "github.com/aws/aws-cdk-go/awsmedialivealpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// MS Smooth input loss action.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import medialive_alpha "github.com/aws/aws-cdk-go/awsmedialivealpha"
//
//   msSmoothInputLossAction := medialive_alpha.MsSmoothInputLossAction_Of(jsii.String("value"))
//
// Experimental.
type MsSmoothInputLossAction interface {
	// The underlying string value passed to CloudFormation.
	// Experimental.
	Value() *string
}

// The jsii proxy struct for MsSmoothInputLossAction
type jsiiProxy_MsSmoothInputLossAction struct {
	_ byte // padding
}

func (j *jsiiProxy_MsSmoothInputLossAction) Value() *string {
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
func MsSmoothInputLossAction_Of(value *string) MsSmoothInputLossAction {
	_init_.Initialize()

	if err := validateMsSmoothInputLossAction_OfParameters(value); err != nil {
		panic(err)
	}
	var returns MsSmoothInputLossAction

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-medialive-alpha.MsSmoothInputLossAction",
		"of",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func MsSmoothInputLossAction_EMIT_OUTPUT() MsSmoothInputLossAction {
	_init_.Initialize()
	var returns MsSmoothInputLossAction
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.MsSmoothInputLossAction",
		"EMIT_OUTPUT",
		&returns,
	)
	return returns
}

func MsSmoothInputLossAction_PAUSE_OUTPUT() MsSmoothInputLossAction {
	_init_.Initialize()
	var returns MsSmoothInputLossAction
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.MsSmoothInputLossAction",
		"PAUSE_OUTPUT",
		&returns,
	)
	return returns
}

