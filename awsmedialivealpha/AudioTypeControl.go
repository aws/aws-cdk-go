package awsmedialivealpha

import (
	_init_ "github.com/aws/aws-cdk-go/awsmedialivealpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Determines how the audio type is signaled in the output.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import medialive_alpha "github.com/aws/aws-cdk-go/awsmedialivealpha"
//
//   audioTypeControl := medialive_alpha.AudioTypeControl_Of(jsii.String("value"))
//
// Experimental.
type AudioTypeControl interface {
	// The underlying string value passed to CloudFormation.
	// Experimental.
	Value() *string
}

// The jsii proxy struct for AudioTypeControl
type jsiiProxy_AudioTypeControl struct {
	_ byte // padding
}

func (j *jsiiProxy_AudioTypeControl) Value() *string {
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
func AudioTypeControl_Of(value *string) AudioTypeControl {
	_init_.Initialize()

	if err := validateAudioTypeControl_OfParameters(value); err != nil {
		panic(err)
	}
	var returns AudioTypeControl

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-medialive-alpha.AudioTypeControl",
		"of",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func AudioTypeControl_FOLLOW_INPUT() AudioTypeControl {
	_init_.Initialize()
	var returns AudioTypeControl
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.AudioTypeControl",
		"FOLLOW_INPUT",
		&returns,
	)
	return returns
}

func AudioTypeControl_USE_CONFIGURED() AudioTypeControl {
	_init_.Initialize()
	var returns AudioTypeControl
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.AudioTypeControl",
		"USE_CONFIGURED",
		&returns,
	)
	return returns
}

