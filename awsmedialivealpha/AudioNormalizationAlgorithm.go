package awsmedialivealpha

import (
	_init_ "github.com/aws/aws-cdk-go/awsmedialivealpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Audio normalization algorithm.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import medialive_alpha "github.com/aws/aws-cdk-go/awsmedialivealpha"
//
//   audioNormalizationAlgorithm := medialive_alpha.AudioNormalizationAlgorithm_Of(jsii.String("value"))
//
// Experimental.
type AudioNormalizationAlgorithm interface {
	// The underlying string value passed to CloudFormation.
	// Experimental.
	Value() *string
}

// The jsii proxy struct for AudioNormalizationAlgorithm
type jsiiProxy_AudioNormalizationAlgorithm struct {
	_ byte // padding
}

func (j *jsiiProxy_AudioNormalizationAlgorithm) Value() *string {
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
func AudioNormalizationAlgorithm_Of(value *string) AudioNormalizationAlgorithm {
	_init_.Initialize()

	if err := validateAudioNormalizationAlgorithm_OfParameters(value); err != nil {
		panic(err)
	}
	var returns AudioNormalizationAlgorithm

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-medialive-alpha.AudioNormalizationAlgorithm",
		"of",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func AudioNormalizationAlgorithm_ITU_1770_1() AudioNormalizationAlgorithm {
	_init_.Initialize()
	var returns AudioNormalizationAlgorithm
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.AudioNormalizationAlgorithm",
		"ITU_1770_1",
		&returns,
	)
	return returns
}

func AudioNormalizationAlgorithm_ITU_1770_2() AudioNormalizationAlgorithm {
	_init_.Initialize()
	var returns AudioNormalizationAlgorithm
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.AudioNormalizationAlgorithm",
		"ITU_1770_2",
		&returns,
	)
	return returns
}

