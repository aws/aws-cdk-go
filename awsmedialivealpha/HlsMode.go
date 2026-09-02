package awsmedialivealpha

import (
	_init_ "github.com/aws/aws-cdk-go/awsmedialivealpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// HLS output mode.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import medialive_alpha "github.com/aws/aws-cdk-go/awsmedialivealpha"
//
//   hlsMode := medialive_alpha.HlsMode_Of(jsii.String("value"))
//
// Experimental.
type HlsMode interface {
	// The underlying string value passed to CloudFormation.
	// Experimental.
	Value() *string
}

// The jsii proxy struct for HlsMode
type jsiiProxy_HlsMode struct {
	_ byte // padding
}

func (j *jsiiProxy_HlsMode) Value() *string {
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
func HlsMode_Of(value *string) HlsMode {
	_init_.Initialize()

	if err := validateHlsMode_OfParameters(value); err != nil {
		panic(err)
	}
	var returns HlsMode

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-medialive-alpha.HlsMode",
		"of",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func HlsMode_LIVE() HlsMode {
	_init_.Initialize()
	var returns HlsMode
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.HlsMode",
		"LIVE",
		&returns,
	)
	return returns
}

func HlsMode_VOD() HlsMode {
	_init_.Initialize()
	var returns HlsMode
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.HlsMode",
		"VOD",
		&returns,
	)
	return returns
}

