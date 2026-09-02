package awsmedialivealpha

import (
	_init_ "github.com/aws/aws-cdk-go/awsmedialivealpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// EAC3 coding mode.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import medialive_alpha "github.com/aws/aws-cdk-go/awsmedialivealpha"
//
//   eac3CodingMode := medialive_alpha.Eac3CodingMode_CODING_MODE_1_0()
//
// Experimental.
type Eac3CodingMode interface {
	// The underlying string value passed to CloudFormation.
	// Experimental.
	Value() *string
}

// The jsii proxy struct for Eac3CodingMode
type jsiiProxy_Eac3CodingMode struct {
	_ byte // padding
}

func (j *jsiiProxy_Eac3CodingMode) Value() *string {
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
func Eac3CodingMode_Of(value *string) Eac3CodingMode {
	_init_.Initialize()

	if err := validateEac3CodingMode_OfParameters(value); err != nil {
		panic(err)
	}
	var returns Eac3CodingMode

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-medialive-alpha.Eac3CodingMode",
		"of",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func Eac3CodingMode_CODING_MODE_1_0() Eac3CodingMode {
	_init_.Initialize()
	var returns Eac3CodingMode
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.Eac3CodingMode",
		"CODING_MODE_1_0",
		&returns,
	)
	return returns
}

func Eac3CodingMode_CODING_MODE_2_0() Eac3CodingMode {
	_init_.Initialize()
	var returns Eac3CodingMode
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.Eac3CodingMode",
		"CODING_MODE_2_0",
		&returns,
	)
	return returns
}

func Eac3CodingMode_CODING_MODE_3_2() Eac3CodingMode {
	_init_.Initialize()
	var returns Eac3CodingMode
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.Eac3CodingMode",
		"CODING_MODE_3_2",
		&returns,
	)
	return returns
}

