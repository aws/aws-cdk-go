package awsmedialivealpha

import (
	_init_ "github.com/aws/aws-cdk-go/awsmedialivealpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Determines how the audio language code is signaled in the output.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import medialive_alpha "github.com/aws/aws-cdk-go/awsmedialivealpha"
//
//   audioLanguageCodeControl := medialive_alpha.AudioLanguageCodeControl_Of(jsii.String("value"))
//
// Experimental.
type AudioLanguageCodeControl interface {
	// The underlying string value passed to CloudFormation.
	// Experimental.
	Value() *string
}

// The jsii proxy struct for AudioLanguageCodeControl
type jsiiProxy_AudioLanguageCodeControl struct {
	_ byte // padding
}

func (j *jsiiProxy_AudioLanguageCodeControl) Value() *string {
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
func AudioLanguageCodeControl_Of(value *string) AudioLanguageCodeControl {
	_init_.Initialize()

	if err := validateAudioLanguageCodeControl_OfParameters(value); err != nil {
		panic(err)
	}
	var returns AudioLanguageCodeControl

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-medialive-alpha.AudioLanguageCodeControl",
		"of",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func AudioLanguageCodeControl_FOLLOW_INPUT() AudioLanguageCodeControl {
	_init_.Initialize()
	var returns AudioLanguageCodeControl
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.AudioLanguageCodeControl",
		"FOLLOW_INPUT",
		&returns,
	)
	return returns
}

func AudioLanguageCodeControl_USE_CONFIGURED() AudioLanguageCodeControl {
	_init_.Initialize()
	var returns AudioLanguageCodeControl
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.AudioLanguageCodeControl",
		"USE_CONFIGURED",
		&returns,
	)
	return returns
}

