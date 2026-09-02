package awsmedialivealpha

import (
	_init_ "github.com/aws/aws-cdk-go/awsmedialivealpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// H.264 adaptive quantization strength.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import medialive_alpha "github.com/aws/aws-cdk-go/awsmedialivealpha"
//
//   h264AdaptiveQuantization := medialive_alpha.H264AdaptiveQuantization_AUTO()
//
// Experimental.
type H264AdaptiveQuantization interface {
	// The underlying string value passed to CloudFormation.
	// Experimental.
	Value() *string
}

// The jsii proxy struct for H264AdaptiveQuantization
type jsiiProxy_H264AdaptiveQuantization struct {
	_ byte // padding
}

func (j *jsiiProxy_H264AdaptiveQuantization) Value() *string {
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
func H264AdaptiveQuantization_Of(value *string) H264AdaptiveQuantization {
	_init_.Initialize()

	if err := validateH264AdaptiveQuantization_OfParameters(value); err != nil {
		panic(err)
	}
	var returns H264AdaptiveQuantization

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-medialive-alpha.H264AdaptiveQuantization",
		"of",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func H264AdaptiveQuantization_AUTO() H264AdaptiveQuantization {
	_init_.Initialize()
	var returns H264AdaptiveQuantization
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.H264AdaptiveQuantization",
		"AUTO",
		&returns,
	)
	return returns
}

func H264AdaptiveQuantization_HIGH() H264AdaptiveQuantization {
	_init_.Initialize()
	var returns H264AdaptiveQuantization
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.H264AdaptiveQuantization",
		"HIGH",
		&returns,
	)
	return returns
}

func H264AdaptiveQuantization_HIGHER() H264AdaptiveQuantization {
	_init_.Initialize()
	var returns H264AdaptiveQuantization
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.H264AdaptiveQuantization",
		"HIGHER",
		&returns,
	)
	return returns
}

func H264AdaptiveQuantization_LOW() H264AdaptiveQuantization {
	_init_.Initialize()
	var returns H264AdaptiveQuantization
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.H264AdaptiveQuantization",
		"LOW",
		&returns,
	)
	return returns
}

func H264AdaptiveQuantization_MAX() H264AdaptiveQuantization {
	_init_.Initialize()
	var returns H264AdaptiveQuantization
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.H264AdaptiveQuantization",
		"MAX",
		&returns,
	)
	return returns
}

func H264AdaptiveQuantization_MEDIUM() H264AdaptiveQuantization {
	_init_.Initialize()
	var returns H264AdaptiveQuantization
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.H264AdaptiveQuantization",
		"MEDIUM",
		&returns,
	)
	return returns
}

func H264AdaptiveQuantization_OFF() H264AdaptiveQuantization {
	_init_.Initialize()
	var returns H264AdaptiveQuantization
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.H264AdaptiveQuantization",
		"OFF",
		&returns,
	)
	return returns
}

