package awsmedialivealpha

import (
	_init_ "github.com/aws/aws-cdk-go/awsmedialivealpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// MS Smooth audio-only timecode control.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import medialive_alpha "github.com/aws/aws-cdk-go/awsmedialivealpha"
//
//   msSmoothAudioOnlyTimecodeControl := medialive_alpha.MsSmoothAudioOnlyTimecodeControl_Of(jsii.String("value"))
//
// Experimental.
type MsSmoothAudioOnlyTimecodeControl interface {
	// The underlying string value passed to CloudFormation.
	// Experimental.
	Value() *string
}

// The jsii proxy struct for MsSmoothAudioOnlyTimecodeControl
type jsiiProxy_MsSmoothAudioOnlyTimecodeControl struct {
	_ byte // padding
}

func (j *jsiiProxy_MsSmoothAudioOnlyTimecodeControl) Value() *string {
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
func MsSmoothAudioOnlyTimecodeControl_Of(value *string) MsSmoothAudioOnlyTimecodeControl {
	_init_.Initialize()

	if err := validateMsSmoothAudioOnlyTimecodeControl_OfParameters(value); err != nil {
		panic(err)
	}
	var returns MsSmoothAudioOnlyTimecodeControl

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-medialive-alpha.MsSmoothAudioOnlyTimecodeControl",
		"of",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func MsSmoothAudioOnlyTimecodeControl_PASSTHROUGH() MsSmoothAudioOnlyTimecodeControl {
	_init_.Initialize()
	var returns MsSmoothAudioOnlyTimecodeControl
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.MsSmoothAudioOnlyTimecodeControl",
		"PASSTHROUGH",
		&returns,
	)
	return returns
}

func MsSmoothAudioOnlyTimecodeControl_USE_CONFIGURED_CLOCK() MsSmoothAudioOnlyTimecodeControl {
	_init_.Initialize()
	var returns MsSmoothAudioOnlyTimecodeControl
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.MsSmoothAudioOnlyTimecodeControl",
		"USE_CONFIGURED_CLOCK",
		&returns,
	)
	return returns
}

