package awsmedialivealpha

import (
	_init_ "github.com/aws/aws-cdk-go/awsmedialivealpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// H.264 syntax mode.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import medialive_alpha "github.com/aws/aws-cdk-go/awsmedialivealpha"
//
//   h264Syntax := medialive_alpha.H264Syntax_Of(jsii.String("value"))
//
// Experimental.
type H264Syntax interface {
	// The underlying string value passed to CloudFormation.
	// Experimental.
	Value() *string
}

// The jsii proxy struct for H264Syntax
type jsiiProxy_H264Syntax struct {
	_ byte // padding
}

func (j *jsiiProxy_H264Syntax) Value() *string {
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
func H264Syntax_Of(value *string) H264Syntax {
	_init_.Initialize()

	if err := validateH264Syntax_OfParameters(value); err != nil {
		panic(err)
	}
	var returns H264Syntax

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-medialive-alpha.H264Syntax",
		"of",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func H264Syntax_DEFAULT() H264Syntax {
	_init_.Initialize()
	var returns H264Syntax
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.H264Syntax",
		"DEFAULT",
		&returns,
	)
	return returns
}

func H264Syntax_RP2027() H264Syntax {
	_init_.Initialize()
	var returns H264Syntax
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.H264Syntax",
		"RP2027",
		&returns,
	)
	return returns
}

