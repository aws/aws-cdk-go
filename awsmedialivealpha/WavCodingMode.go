package awsmedialivealpha

import (
	_init_ "github.com/aws/aws-cdk-go/awsmedialivealpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// WAV coding mode.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import medialive_alpha "github.com/aws/aws-cdk-go/awsmedialivealpha"
//
//   wavCodingMode := medialive_alpha.WavCodingMode_CODING_MODE_1_0()
//
// Experimental.
type WavCodingMode interface {
	// The underlying string value passed to CloudFormation.
	// Experimental.
	Value() *string
}

// The jsii proxy struct for WavCodingMode
type jsiiProxy_WavCodingMode struct {
	_ byte // padding
}

func (j *jsiiProxy_WavCodingMode) Value() *string {
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
func WavCodingMode_Of(value *string) WavCodingMode {
	_init_.Initialize()

	if err := validateWavCodingMode_OfParameters(value); err != nil {
		panic(err)
	}
	var returns WavCodingMode

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-medialive-alpha.WavCodingMode",
		"of",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func WavCodingMode_CODING_MODE_1_0() WavCodingMode {
	_init_.Initialize()
	var returns WavCodingMode
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.WavCodingMode",
		"CODING_MODE_1_0",
		&returns,
	)
	return returns
}

func WavCodingMode_CODING_MODE_2_0() WavCodingMode {
	_init_.Initialize()
	var returns WavCodingMode
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.WavCodingMode",
		"CODING_MODE_2_0",
		&returns,
	)
	return returns
}

func WavCodingMode_CODING_MODE_4_0() WavCodingMode {
	_init_.Initialize()
	var returns WavCodingMode
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.WavCodingMode",
		"CODING_MODE_4_0",
		&returns,
	)
	return returns
}

func WavCodingMode_CODING_MODE_8_0() WavCodingMode {
	_init_.Initialize()
	var returns WavCodingMode
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.WavCodingMode",
		"CODING_MODE_8_0",
		&returns,
	)
	return returns
}

