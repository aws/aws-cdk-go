package awsmedialivealpha

import (
	_init_ "github.com/aws/aws-cdk-go/awsmedialivealpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Peak calculation method for audio normalization.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import medialive_alpha "github.com/aws/aws-cdk-go/awsmedialivealpha"
//
//   audioNormalizationPeakCalculation := medialive_alpha.AudioNormalizationPeakCalculation_Of(jsii.String("value"))
//
// Experimental.
type AudioNormalizationPeakCalculation interface {
	// The underlying string value passed to CloudFormation.
	// Experimental.
	Value() *string
}

// The jsii proxy struct for AudioNormalizationPeakCalculation
type jsiiProxy_AudioNormalizationPeakCalculation struct {
	_ byte // padding
}

func (j *jsiiProxy_AudioNormalizationPeakCalculation) Value() *string {
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
func AudioNormalizationPeakCalculation_Of(value *string) AudioNormalizationPeakCalculation {
	_init_.Initialize()

	if err := validateAudioNormalizationPeakCalculation_OfParameters(value); err != nil {
		panic(err)
	}
	var returns AudioNormalizationPeakCalculation

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-medialive-alpha.AudioNormalizationPeakCalculation",
		"of",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func AudioNormalizationPeakCalculation_TRUE_PEAK() AudioNormalizationPeakCalculation {
	_init_.Initialize()
	var returns AudioNormalizationPeakCalculation
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.AudioNormalizationPeakCalculation",
		"TRUE_PEAK",
		&returns,
	)
	return returns
}

