package awsmedialivealpha

import (
	_init_ "github.com/aws/aws-cdk-go/awsmedialivealpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Behavior of last resort when input video is lost and no more backup inputs are available, for an SRT output group.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import medialive_alpha "github.com/aws/aws-cdk-go/awsmedialivealpha"
//
//   srtInputLossAction := medialive_alpha.SrtInputLossAction_DROP_PROGRAM()
//
// Experimental.
type SrtInputLossAction interface {
	// The underlying string value passed to CloudFormation.
	// Experimental.
	Value() *string
}

// The jsii proxy struct for SrtInputLossAction
type jsiiProxy_SrtInputLossAction struct {
	_ byte // padding
}

func (j *jsiiProxy_SrtInputLossAction) Value() *string {
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
func SrtInputLossAction_Of(value *string) SrtInputLossAction {
	_init_.Initialize()

	if err := validateSrtInputLossAction_OfParameters(value); err != nil {
		panic(err)
	}
	var returns SrtInputLossAction

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-medialive-alpha.SrtInputLossAction",
		"of",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func SrtInputLossAction_DROP_PROGRAM() SrtInputLossAction {
	_init_.Initialize()
	var returns SrtInputLossAction
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.SrtInputLossAction",
		"DROP_PROGRAM",
		&returns,
	)
	return returns
}

func SrtInputLossAction_DROP_TS() SrtInputLossAction {
	_init_.Initialize()
	var returns SrtInputLossAction
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.SrtInputLossAction",
		"DROP_TS",
		&returns,
	)
	return returns
}

func SrtInputLossAction_EMIT_PROGRAM() SrtInputLossAction {
	_init_.Initialize()
	var returns SrtInputLossAction
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.SrtInputLossAction",
		"EMIT_PROGRAM",
		&returns,
	)
	return returns
}

