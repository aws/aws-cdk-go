package awsmedialivealpha

import (
	_init_ "github.com/aws/aws-cdk-go/awsmedialivealpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Audio normalization algorithm control.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import medialive_alpha "github.com/aws/aws-cdk-go/awsmedialivealpha"
//
//   audioNormalizationAlgorithmControl := medialive_alpha.AudioNormalizationAlgorithmControl_Of(jsii.String("value"))
//
// Experimental.
type AudioNormalizationAlgorithmControl interface {
	// The underlying string value passed to CloudFormation.
	// Experimental.
	Value() *string
}

// The jsii proxy struct for AudioNormalizationAlgorithmControl
type jsiiProxy_AudioNormalizationAlgorithmControl struct {
	_ byte // padding
}

func (j *jsiiProxy_AudioNormalizationAlgorithmControl) Value() *string {
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
func AudioNormalizationAlgorithmControl_Of(value *string) AudioNormalizationAlgorithmControl {
	_init_.Initialize()

	if err := validateAudioNormalizationAlgorithmControl_OfParameters(value); err != nil {
		panic(err)
	}
	var returns AudioNormalizationAlgorithmControl

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-medialive-alpha.AudioNormalizationAlgorithmControl",
		"of",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func AudioNormalizationAlgorithmControl_CORRECT_AUDIO() AudioNormalizationAlgorithmControl {
	_init_.Initialize()
	var returns AudioNormalizationAlgorithmControl
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.AudioNormalizationAlgorithmControl",
		"CORRECT_AUDIO",
		&returns,
	)
	return returns
}

func AudioNormalizationAlgorithmControl_MEASURE_ONLY() AudioNormalizationAlgorithmControl {
	_init_.Initialize()
	var returns AudioNormalizationAlgorithmControl
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.AudioNormalizationAlgorithmControl",
		"MEASURE_ONLY",
		&returns,
	)
	return returns
}

