package awsmedialivealpha

import (
	_init_ "github.com/aws/aws-cdk-go/awsmedialivealpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// MS Smooth segmentation mode.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import medialive_alpha "github.com/aws/aws-cdk-go/awsmedialivealpha"
//
//   msSmoothSegmentationMode := medialive_alpha.MsSmoothSegmentationMode_Of(jsii.String("value"))
//
// Experimental.
type MsSmoothSegmentationMode interface {
	// The underlying string value passed to CloudFormation.
	// Experimental.
	Value() *string
}

// The jsii proxy struct for MsSmoothSegmentationMode
type jsiiProxy_MsSmoothSegmentationMode struct {
	_ byte // padding
}

func (j *jsiiProxy_MsSmoothSegmentationMode) Value() *string {
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
func MsSmoothSegmentationMode_Of(value *string) MsSmoothSegmentationMode {
	_init_.Initialize()

	if err := validateMsSmoothSegmentationMode_OfParameters(value); err != nil {
		panic(err)
	}
	var returns MsSmoothSegmentationMode

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-medialive-alpha.MsSmoothSegmentationMode",
		"of",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func MsSmoothSegmentationMode_USE_INPUT_SEGMENTATION() MsSmoothSegmentationMode {
	_init_.Initialize()
	var returns MsSmoothSegmentationMode
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.MsSmoothSegmentationMode",
		"USE_INPUT_SEGMENTATION",
		&returns,
	)
	return returns
}

func MsSmoothSegmentationMode_USE_SEGMENT_DURATION() MsSmoothSegmentationMode {
	_init_.Initialize()
	var returns MsSmoothSegmentationMode
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.MsSmoothSegmentationMode",
		"USE_SEGMENT_DURATION",
		&returns,
	)
	return returns
}

