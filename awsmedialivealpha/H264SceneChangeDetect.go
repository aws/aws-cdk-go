package awsmedialivealpha

import (
	_init_ "github.com/aws/aws-cdk-go/awsmedialivealpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// H.264 scene change detection.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import medialive_alpha "github.com/aws/aws-cdk-go/awsmedialivealpha"
//
//   h264SceneChangeDetect := medialive_alpha.H264SceneChangeDetect_Of(jsii.String("value"))
//
// Experimental.
type H264SceneChangeDetect interface {
	// The underlying string value passed to CloudFormation.
	// Experimental.
	Value() *string
}

// The jsii proxy struct for H264SceneChangeDetect
type jsiiProxy_H264SceneChangeDetect struct {
	_ byte // padding
}

func (j *jsiiProxy_H264SceneChangeDetect) Value() *string {
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
func H264SceneChangeDetect_Of(value *string) H264SceneChangeDetect {
	_init_.Initialize()

	if err := validateH264SceneChangeDetect_OfParameters(value); err != nil {
		panic(err)
	}
	var returns H264SceneChangeDetect

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-medialive-alpha.H264SceneChangeDetect",
		"of",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func H264SceneChangeDetect_DISABLED() H264SceneChangeDetect {
	_init_.Initialize()
	var returns H264SceneChangeDetect
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.H264SceneChangeDetect",
		"DISABLED",
		&returns,
	)
	return returns
}

func H264SceneChangeDetect_ENABLED() H264SceneChangeDetect {
	_init_.Initialize()
	var returns H264SceneChangeDetect
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.H264SceneChangeDetect",
		"ENABLED",
		&returns,
	)
	return returns
}

