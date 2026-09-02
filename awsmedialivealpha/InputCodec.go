package awsmedialivealpha

import (
	_init_ "github.com/aws/aws-cdk-go/awsmedialivealpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// The codec for the input specification.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import medialive_alpha "github.com/aws/aws-cdk-go/awsmedialivealpha"
//
//   inputCodec := medialive_alpha.InputCodec_AVC()
//
// Experimental.
type InputCodec interface {
	// The underlying string value passed to CloudFormation.
	// Experimental.
	Value() *string
}

// The jsii proxy struct for InputCodec
type jsiiProxy_InputCodec struct {
	_ byte // padding
}

func (j *jsiiProxy_InputCodec) Value() *string {
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
func InputCodec_Of(value *string) InputCodec {
	_init_.Initialize()

	if err := validateInputCodec_OfParameters(value); err != nil {
		panic(err)
	}
	var returns InputCodec

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-medialive-alpha.InputCodec",
		"of",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func InputCodec_AVC() InputCodec {
	_init_.Initialize()
	var returns InputCodec
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.InputCodec",
		"AVC",
		&returns,
	)
	return returns
}

func InputCodec_HEVC() InputCodec {
	_init_.Initialize()
	var returns InputCodec
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.InputCodec",
		"HEVC",
		&returns,
	)
	return returns
}

func InputCodec_MPEG2() InputCodec {
	_init_.Initialize()
	var returns InputCodec
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.InputCodec",
		"MPEG2",
		&returns,
	)
	return returns
}

