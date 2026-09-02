package awsmedialivealpha

import (
	_init_ "github.com/aws/aws-cdk-go/awsmedialivealpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// HLS segmentation mode.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import medialive_alpha "github.com/aws/aws-cdk-go/awsmedialivealpha"
//
//   hlsSegmentationMode := medialive_alpha.HlsSegmentationMode_Of(jsii.String("value"))
//
// Experimental.
type HlsSegmentationMode interface {
	// The underlying string value passed to CloudFormation.
	// Experimental.
	Value() *string
}

// The jsii proxy struct for HlsSegmentationMode
type jsiiProxy_HlsSegmentationMode struct {
	_ byte // padding
}

func (j *jsiiProxy_HlsSegmentationMode) Value() *string {
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
func HlsSegmentationMode_Of(value *string) HlsSegmentationMode {
	_init_.Initialize()

	if err := validateHlsSegmentationMode_OfParameters(value); err != nil {
		panic(err)
	}
	var returns HlsSegmentationMode

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-medialive-alpha.HlsSegmentationMode",
		"of",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func HlsSegmentationMode_USE_INPUT_SEGMENTATION() HlsSegmentationMode {
	_init_.Initialize()
	var returns HlsSegmentationMode
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.HlsSegmentationMode",
		"USE_INPUT_SEGMENTATION",
		&returns,
	)
	return returns
}

func HlsSegmentationMode_USE_SEGMENT_DURATION() HlsSegmentationMode {
	_init_.Initialize()
	var returns HlsSegmentationMode
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.HlsSegmentationMode",
		"USE_SEGMENT_DURATION",
		&returns,
	)
	return returns
}

