package awsmedialivealpha

import (
	_init_ "github.com/aws/aws-cdk-go/awsmedialivealpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Controls whether MediaLive delays video to synchronize captions with audio and video output.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import medialive_alpha "github.com/aws/aws-cdk-go/awsmedialivealpha"
//
//   captionSynchronizationMode := medialive_alpha.CaptionSynchronizationMode_Of(jsii.String("value"))
//
// Experimental.
type CaptionSynchronizationMode interface {
	// The underlying string value passed to CloudFormation.
	// Experimental.
	Value() *string
}

// The jsii proxy struct for CaptionSynchronizationMode
type jsiiProxy_CaptionSynchronizationMode struct {
	_ byte // padding
}

func (j *jsiiProxy_CaptionSynchronizationMode) Value() *string {
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
func CaptionSynchronizationMode_Of(value *string) CaptionSynchronizationMode {
	_init_.Initialize()

	if err := validateCaptionSynchronizationMode_OfParameters(value); err != nil {
		panic(err)
	}
	var returns CaptionSynchronizationMode

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-medialive-alpha.CaptionSynchronizationMode",
		"of",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func CaptionSynchronizationMode_NO_VIDEO_DELAY() CaptionSynchronizationMode {
	_init_.Initialize()
	var returns CaptionSynchronizationMode
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.CaptionSynchronizationMode",
		"NO_VIDEO_DELAY",
		&returns,
	)
	return returns
}

func CaptionSynchronizationMode_VIDEO_ALIGNED_CAPTIONS() CaptionSynchronizationMode {
	_init_.Initialize()
	var returns CaptionSynchronizationMode
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.CaptionSynchronizationMode",
		"VIDEO_ALIGNED_CAPTIONS",
		&returns,
	)
	return returns
}

