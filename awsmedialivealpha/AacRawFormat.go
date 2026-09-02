package awsmedialivealpha

import (
	_init_ "github.com/aws/aws-cdk-go/awsmedialivealpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// AAC raw format.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import medialive_alpha "github.com/aws/aws-cdk-go/awsmedialivealpha"
//
//   aacRawFormat := medialive_alpha.AacRawFormat_Of(jsii.String("value"))
//
// Experimental.
type AacRawFormat interface {
	// The underlying string value passed to CloudFormation.
	// Experimental.
	Value() *string
}

// The jsii proxy struct for AacRawFormat
type jsiiProxy_AacRawFormat struct {
	_ byte // padding
}

func (j *jsiiProxy_AacRawFormat) Value() *string {
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
func AacRawFormat_Of(value *string) AacRawFormat {
	_init_.Initialize()

	if err := validateAacRawFormat_OfParameters(value); err != nil {
		panic(err)
	}
	var returns AacRawFormat

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-medialive-alpha.AacRawFormat",
		"of",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func AacRawFormat_LATM_LOAS() AacRawFormat {
	_init_.Initialize()
	var returns AacRawFormat
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.AacRawFormat",
		"LATM_LOAS",
		&returns,
	)
	return returns
}

func AacRawFormat_NONE() AacRawFormat {
	_init_.Initialize()
	var returns AacRawFormat
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.AacRawFormat",
		"NONE",
		&returns,
	)
	return returns
}

