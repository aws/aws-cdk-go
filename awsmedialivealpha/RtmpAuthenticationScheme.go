package awsmedialivealpha

import (
	_init_ "github.com/aws/aws-cdk-go/awsmedialivealpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// RTMP authentication scheme.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import medialive_alpha "github.com/aws/aws-cdk-go/awsmedialivealpha"
//
//   rtmpAuthenticationScheme := medialive_alpha.RtmpAuthenticationScheme_Of(jsii.String("value"))
//
// Experimental.
type RtmpAuthenticationScheme interface {
	// The underlying string value passed to CloudFormation.
	// Experimental.
	Value() *string
}

// The jsii proxy struct for RtmpAuthenticationScheme
type jsiiProxy_RtmpAuthenticationScheme struct {
	_ byte // padding
}

func (j *jsiiProxy_RtmpAuthenticationScheme) Value() *string {
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
func RtmpAuthenticationScheme_Of(value *string) RtmpAuthenticationScheme {
	_init_.Initialize()

	if err := validateRtmpAuthenticationScheme_OfParameters(value); err != nil {
		panic(err)
	}
	var returns RtmpAuthenticationScheme

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-medialive-alpha.RtmpAuthenticationScheme",
		"of",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func RtmpAuthenticationScheme_AKAMAI() RtmpAuthenticationScheme {
	_init_.Initialize()
	var returns RtmpAuthenticationScheme
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.RtmpAuthenticationScheme",
		"AKAMAI",
		&returns,
	)
	return returns
}

func RtmpAuthenticationScheme_COMMON() RtmpAuthenticationScheme {
	_init_.Initialize()
	var returns RtmpAuthenticationScheme
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.RtmpAuthenticationScheme",
		"COMMON",
		&returns,
	)
	return returns
}

