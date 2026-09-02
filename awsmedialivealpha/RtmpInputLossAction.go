package awsmedialivealpha

import (
	_init_ "github.com/aws/aws-cdk-go/awsmedialivealpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// RTMP input loss action.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import medialive_alpha "github.com/aws/aws-cdk-go/awsmedialivealpha"
//
//   rtmpInputLossAction := medialive_alpha.RtmpInputLossAction_Of(jsii.String("value"))
//
// Experimental.
type RtmpInputLossAction interface {
	// The underlying string value passed to CloudFormation.
	// Experimental.
	Value() *string
}

// The jsii proxy struct for RtmpInputLossAction
type jsiiProxy_RtmpInputLossAction struct {
	_ byte // padding
}

func (j *jsiiProxy_RtmpInputLossAction) Value() *string {
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
func RtmpInputLossAction_Of(value *string) RtmpInputLossAction {
	_init_.Initialize()

	if err := validateRtmpInputLossAction_OfParameters(value); err != nil {
		panic(err)
	}
	var returns RtmpInputLossAction

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-medialive-alpha.RtmpInputLossAction",
		"of",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func RtmpInputLossAction_EMIT_OUTPUT() RtmpInputLossAction {
	_init_.Initialize()
	var returns RtmpInputLossAction
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.RtmpInputLossAction",
		"EMIT_OUTPUT",
		&returns,
	)
	return returns
}

func RtmpInputLossAction_PAUSE_OUTPUT() RtmpInputLossAction {
	_init_.Initialize()
	var returns RtmpInputLossAction
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.RtmpInputLossAction",
		"PAUSE_OUTPUT",
		&returns,
	)
	return returns
}

