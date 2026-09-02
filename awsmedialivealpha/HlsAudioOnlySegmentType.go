package awsmedialivealpha

import (
	_init_ "github.com/aws/aws-cdk-go/awsmedialivealpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// The segment container type for an audio-only HLS output.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import medialive_alpha "github.com/aws/aws-cdk-go/awsmedialivealpha"
//
//   hlsAudioOnlySegmentType := medialive_alpha.HlsAudioOnlySegmentType_Of(jsii.String("value"))
//
// Experimental.
type HlsAudioOnlySegmentType interface {
	// The underlying string value passed to CloudFormation.
	// Experimental.
	Value() *string
}

// The jsii proxy struct for HlsAudioOnlySegmentType
type jsiiProxy_HlsAudioOnlySegmentType struct {
	_ byte // padding
}

func (j *jsiiProxy_HlsAudioOnlySegmentType) Value() *string {
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
func HlsAudioOnlySegmentType_Of(value *string) HlsAudioOnlySegmentType {
	_init_.Initialize()

	if err := validateHlsAudioOnlySegmentType_OfParameters(value); err != nil {
		panic(err)
	}
	var returns HlsAudioOnlySegmentType

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-medialive-alpha.HlsAudioOnlySegmentType",
		"of",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func HlsAudioOnlySegmentType_AAC() HlsAudioOnlySegmentType {
	_init_.Initialize()
	var returns HlsAudioOnlySegmentType
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.HlsAudioOnlySegmentType",
		"AAC",
		&returns,
	)
	return returns
}

func HlsAudioOnlySegmentType_FMP4() HlsAudioOnlySegmentType {
	_init_.Initialize()
	var returns HlsAudioOnlySegmentType
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.HlsAudioOnlySegmentType",
		"FMP4",
		&returns,
	)
	return returns
}

