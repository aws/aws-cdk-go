package awsmedialivealpha

import (
	_init_ "github.com/aws/aws-cdk-go/awsmedialivealpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// H.264 spatial adaptive quantization.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import medialive_alpha "github.com/aws/aws-cdk-go/awsmedialivealpha"
//
//   h264SpatialAq := medialive_alpha.H264SpatialAq_Of(jsii.String("value"))
//
// Experimental.
type H264SpatialAq interface {
	// The underlying string value passed to CloudFormation.
	// Experimental.
	Value() *string
}

// The jsii proxy struct for H264SpatialAq
type jsiiProxy_H264SpatialAq struct {
	_ byte // padding
}

func (j *jsiiProxy_H264SpatialAq) Value() *string {
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
func H264SpatialAq_Of(value *string) H264SpatialAq {
	_init_.Initialize()

	if err := validateH264SpatialAq_OfParameters(value); err != nil {
		panic(err)
	}
	var returns H264SpatialAq

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-medialive-alpha.H264SpatialAq",
		"of",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func H264SpatialAq_DISABLED() H264SpatialAq {
	_init_.Initialize()
	var returns H264SpatialAq
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.H264SpatialAq",
		"DISABLED",
		&returns,
	)
	return returns
}

func H264SpatialAq_ENABLED() H264SpatialAq {
	_init_.Initialize()
	var returns H264SpatialAq
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.H264SpatialAq",
		"ENABLED",
		&returns,
	)
	return returns
}

