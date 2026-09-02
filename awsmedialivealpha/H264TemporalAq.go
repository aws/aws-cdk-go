package awsmedialivealpha

import (
	_init_ "github.com/aws/aws-cdk-go/awsmedialivealpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// H.264 temporal adaptive quantization.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import medialive_alpha "github.com/aws/aws-cdk-go/awsmedialivealpha"
//
//   h264TemporalAq := medialive_alpha.H264TemporalAq_Of(jsii.String("value"))
//
// Experimental.
type H264TemporalAq interface {
	// The underlying string value passed to CloudFormation.
	// Experimental.
	Value() *string
}

// The jsii proxy struct for H264TemporalAq
type jsiiProxy_H264TemporalAq struct {
	_ byte // padding
}

func (j *jsiiProxy_H264TemporalAq) Value() *string {
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
func H264TemporalAq_Of(value *string) H264TemporalAq {
	_init_.Initialize()

	if err := validateH264TemporalAq_OfParameters(value); err != nil {
		panic(err)
	}
	var returns H264TemporalAq

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-medialive-alpha.H264TemporalAq",
		"of",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func H264TemporalAq_DISABLED() H264TemporalAq {
	_init_.Initialize()
	var returns H264TemporalAq
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.H264TemporalAq",
		"DISABLED",
		&returns,
	)
	return returns
}

func H264TemporalAq_ENABLED() H264TemporalAq {
	_init_.Initialize()
	var returns H264TemporalAq
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.H264TemporalAq",
		"ENABLED",
		&returns,
	)
	return returns
}

