package awsmedialivealpha

import (
	_init_ "github.com/aws/aws-cdk-go/awsmedialivealpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// HLS IV source.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import medialive_alpha "github.com/aws/aws-cdk-go/awsmedialivealpha"
//
//   hlsIvSource := medialive_alpha.HlsIvSource_Of(jsii.String("value"))
//
// Experimental.
type HlsIvSource interface {
	// The underlying string value passed to CloudFormation.
	// Experimental.
	Value() *string
}

// The jsii proxy struct for HlsIvSource
type jsiiProxy_HlsIvSource struct {
	_ byte // padding
}

func (j *jsiiProxy_HlsIvSource) Value() *string {
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
func HlsIvSource_Of(value *string) HlsIvSource {
	_init_.Initialize()

	if err := validateHlsIvSource_OfParameters(value); err != nil {
		panic(err)
	}
	var returns HlsIvSource

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-medialive-alpha.HlsIvSource",
		"of",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func HlsIvSource_EXPLICIT() HlsIvSource {
	_init_.Initialize()
	var returns HlsIvSource
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.HlsIvSource",
		"EXPLICIT",
		&returns,
	)
	return returns
}

func HlsIvSource_FOLLOWS_SEGMENT_NUMBER() HlsIvSource {
	_init_.Initialize()
	var returns HlsIvSource
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.HlsIvSource",
		"FOLLOWS_SEGMENT_NUMBER",
		&returns,
	)
	return returns
}

