package awsmedialivealpha

import (
	_init_ "github.com/aws/aws-cdk-go/awsmedialivealpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Audio bit depth for WAV codec.
//
// Use one of the standard presets or `AudioBitDepth.of(bits)` for a custom value.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import medialive_alpha "github.com/aws/aws-cdk-go/awsmedialivealpha"
//
//   audioBitDepth := medialive_alpha.AudioBitDepth_Of(jsii.Number(123))
//
// Experimental.
type AudioBitDepth interface {
}

// The jsii proxy struct for AudioBitDepth
type jsiiProxy_AudioBitDepth struct {
	_ byte // padding
}

// A custom bit depth.
// Experimental.
func AudioBitDepth_Of(bits *float64) AudioBitDepth {
	_init_.Initialize()

	if err := validateAudioBitDepth_OfParameters(bits); err != nil {
		panic(err)
	}
	var returns AudioBitDepth

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-medialive-alpha.AudioBitDepth",
		"of",
		[]interface{}{bits},
		&returns,
	)

	return returns
}

func AudioBitDepth_DEPTH_16() AudioBitDepth {
	_init_.Initialize()
	var returns AudioBitDepth
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.AudioBitDepth",
		"DEPTH_16",
		&returns,
	)
	return returns
}

func AudioBitDepth_DEPTH_24() AudioBitDepth {
	_init_.Initialize()
	var returns AudioBitDepth
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.AudioBitDepth",
		"DEPTH_24",
		&returns,
	)
	return returns
}

