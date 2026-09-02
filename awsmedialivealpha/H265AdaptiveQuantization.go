package awsmedialivealpha

import (
	_init_ "github.com/aws/aws-cdk-go/awsmedialivealpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// H.265 adaptive quantization.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import medialive_alpha "github.com/aws/aws-cdk-go/awsmedialivealpha"
//
//   h265AdaptiveQuantization := medialive_alpha.H265AdaptiveQuantization_AUTO()
//
// Experimental.
type H265AdaptiveQuantization interface {
	// The underlying string value passed to CloudFormation.
	// Experimental.
	Value() *string
}

// The jsii proxy struct for H265AdaptiveQuantization
type jsiiProxy_H265AdaptiveQuantization struct {
	_ byte // padding
}

func (j *jsiiProxy_H265AdaptiveQuantization) Value() *string {
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
func H265AdaptiveQuantization_Of(value *string) H265AdaptiveQuantization {
	_init_.Initialize()

	if err := validateH265AdaptiveQuantization_OfParameters(value); err != nil {
		panic(err)
	}
	var returns H265AdaptiveQuantization

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-medialive-alpha.H265AdaptiveQuantization",
		"of",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func H265AdaptiveQuantization_AUTO() H265AdaptiveQuantization {
	_init_.Initialize()
	var returns H265AdaptiveQuantization
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.H265AdaptiveQuantization",
		"AUTO",
		&returns,
	)
	return returns
}

func H265AdaptiveQuantization_HIGH() H265AdaptiveQuantization {
	_init_.Initialize()
	var returns H265AdaptiveQuantization
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.H265AdaptiveQuantization",
		"HIGH",
		&returns,
	)
	return returns
}

func H265AdaptiveQuantization_HIGHER() H265AdaptiveQuantization {
	_init_.Initialize()
	var returns H265AdaptiveQuantization
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.H265AdaptiveQuantization",
		"HIGHER",
		&returns,
	)
	return returns
}

func H265AdaptiveQuantization_LOW() H265AdaptiveQuantization {
	_init_.Initialize()
	var returns H265AdaptiveQuantization
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.H265AdaptiveQuantization",
		"LOW",
		&returns,
	)
	return returns
}

func H265AdaptiveQuantization_MAX() H265AdaptiveQuantization {
	_init_.Initialize()
	var returns H265AdaptiveQuantization
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.H265AdaptiveQuantization",
		"MAX",
		&returns,
	)
	return returns
}

func H265AdaptiveQuantization_MEDIUM() H265AdaptiveQuantization {
	_init_.Initialize()
	var returns H265AdaptiveQuantization
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.H265AdaptiveQuantization",
		"MEDIUM",
		&returns,
	)
	return returns
}

func H265AdaptiveQuantization_OFF() H265AdaptiveQuantization {
	_init_.Initialize()
	var returns H265AdaptiveQuantization
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.H265AdaptiveQuantization",
		"OFF",
		&returns,
	)
	return returns
}

