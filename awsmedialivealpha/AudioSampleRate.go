package awsmedialivealpha

import (
	_init_ "github.com/aws/aws-cdk-go/awsmedialivealpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Audio sample rate for AAC, MP2, and WAV codecs.
//
// Use one of the standard presets or `AudioSampleRate.of(hz)` for a custom value.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import medialive_alpha "github.com/aws/aws-cdk-go/awsmedialivealpha"
//
//   audioSampleRate := medialive_alpha.AudioSampleRate_HZ_12000()
//
// Experimental.
type AudioSampleRate interface {
}

// The jsii proxy struct for AudioSampleRate
type jsiiProxy_AudioSampleRate struct {
	_ byte // padding
}

// A custom sample rate in Hz.
// Experimental.
func AudioSampleRate_Of(hz *float64) AudioSampleRate {
	_init_.Initialize()

	if err := validateAudioSampleRate_OfParameters(hz); err != nil {
		panic(err)
	}
	var returns AudioSampleRate

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-medialive-alpha.AudioSampleRate",
		"of",
		[]interface{}{hz},
		&returns,
	)

	return returns
}

func AudioSampleRate_HZ_12000() AudioSampleRate {
	_init_.Initialize()
	var returns AudioSampleRate
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.AudioSampleRate",
		"HZ_12000",
		&returns,
	)
	return returns
}

func AudioSampleRate_HZ_16000() AudioSampleRate {
	_init_.Initialize()
	var returns AudioSampleRate
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.AudioSampleRate",
		"HZ_16000",
		&returns,
	)
	return returns
}

func AudioSampleRate_HZ_22050() AudioSampleRate {
	_init_.Initialize()
	var returns AudioSampleRate
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.AudioSampleRate",
		"HZ_22050",
		&returns,
	)
	return returns
}

func AudioSampleRate_HZ_24000() AudioSampleRate {
	_init_.Initialize()
	var returns AudioSampleRate
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.AudioSampleRate",
		"HZ_24000",
		&returns,
	)
	return returns
}

func AudioSampleRate_HZ_32000() AudioSampleRate {
	_init_.Initialize()
	var returns AudioSampleRate
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.AudioSampleRate",
		"HZ_32000",
		&returns,
	)
	return returns
}

func AudioSampleRate_HZ_44100() AudioSampleRate {
	_init_.Initialize()
	var returns AudioSampleRate
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.AudioSampleRate",
		"HZ_44100",
		&returns,
	)
	return returns
}

func AudioSampleRate_HZ_48000() AudioSampleRate {
	_init_.Initialize()
	var returns AudioSampleRate
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.AudioSampleRate",
		"HZ_48000",
		&returns,
	)
	return returns
}

func AudioSampleRate_HZ_8000() AudioSampleRate {
	_init_.Initialize()
	var returns AudioSampleRate
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.AudioSampleRate",
		"HZ_8000",
		&returns,
	)
	return returns
}

func AudioSampleRate_HZ_88200() AudioSampleRate {
	_init_.Initialize()
	var returns AudioSampleRate
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.AudioSampleRate",
		"HZ_88200",
		&returns,
	)
	return returns
}

func AudioSampleRate_HZ_96000() AudioSampleRate {
	_init_.Initialize()
	var returns AudioSampleRate
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.AudioSampleRate",
		"HZ_96000",
		&returns,
	)
	return returns
}

