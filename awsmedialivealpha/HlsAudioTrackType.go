package awsmedialivealpha

import (
	_init_ "github.com/aws/aws-cdk-go/awsmedialivealpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// The audio track type for an audio-only HLS output.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import medialive_alpha "github.com/aws/aws-cdk-go/awsmedialivealpha"
//
//   hlsAudioTrackType := medialive_alpha.HlsAudioTrackType_ALTERNATE_AUDIO_AUTO_SELECT()
//
// Experimental.
type HlsAudioTrackType interface {
	// The underlying string value passed to CloudFormation.
	// Experimental.
	Value() *string
}

// The jsii proxy struct for HlsAudioTrackType
type jsiiProxy_HlsAudioTrackType struct {
	_ byte // padding
}

func (j *jsiiProxy_HlsAudioTrackType) Value() *string {
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
func HlsAudioTrackType_Of(value *string) HlsAudioTrackType {
	_init_.Initialize()

	if err := validateHlsAudioTrackType_OfParameters(value); err != nil {
		panic(err)
	}
	var returns HlsAudioTrackType

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-medialive-alpha.HlsAudioTrackType",
		"of",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func HlsAudioTrackType_ALTERNATE_AUDIO_AUTO_SELECT() HlsAudioTrackType {
	_init_.Initialize()
	var returns HlsAudioTrackType
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.HlsAudioTrackType",
		"ALTERNATE_AUDIO_AUTO_SELECT",
		&returns,
	)
	return returns
}

func HlsAudioTrackType_ALTERNATE_AUDIO_AUTO_SELECT_DEFAULT() HlsAudioTrackType {
	_init_.Initialize()
	var returns HlsAudioTrackType
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.HlsAudioTrackType",
		"ALTERNATE_AUDIO_AUTO_SELECT_DEFAULT",
		&returns,
	)
	return returns
}

func HlsAudioTrackType_ALTERNATE_AUDIO_NOT_AUTO_SELECT() HlsAudioTrackType {
	_init_.Initialize()
	var returns HlsAudioTrackType
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.HlsAudioTrackType",
		"ALTERNATE_AUDIO_NOT_AUTO_SELECT",
		&returns,
	)
	return returns
}

func HlsAudioTrackType_AUDIO_ONLY_VARIANT_STREAM() HlsAudioTrackType {
	_init_.Initialize()
	var returns HlsAudioTrackType
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.HlsAudioTrackType",
		"AUDIO_ONLY_VARIANT_STREAM",
		&returns,
	)
	return returns
}

