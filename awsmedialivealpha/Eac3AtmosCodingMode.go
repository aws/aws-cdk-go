package awsmedialivealpha

import (
	_init_ "github.com/aws/aws-cdk-go/awsmedialivealpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// EAC3 Atmos coding mode.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import medialive_alpha "github.com/aws/aws-cdk-go/awsmedialivealpha"
//
//   eac3AtmosCodingMode := medialive_alpha.Eac3AtmosCodingMode_CODING_MODE_5_1_4()
//
// Experimental.
type Eac3AtmosCodingMode interface {
	// The underlying string value passed to CloudFormation.
	// Experimental.
	Value() *string
}

// The jsii proxy struct for Eac3AtmosCodingMode
type jsiiProxy_Eac3AtmosCodingMode struct {
	_ byte // padding
}

func (j *jsiiProxy_Eac3AtmosCodingMode) Value() *string {
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
func Eac3AtmosCodingMode_Of(value *string) Eac3AtmosCodingMode {
	_init_.Initialize()

	if err := validateEac3AtmosCodingMode_OfParameters(value); err != nil {
		panic(err)
	}
	var returns Eac3AtmosCodingMode

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-medialive-alpha.Eac3AtmosCodingMode",
		"of",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func Eac3AtmosCodingMode_CODING_MODE_5_1_4() Eac3AtmosCodingMode {
	_init_.Initialize()
	var returns Eac3AtmosCodingMode
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.Eac3AtmosCodingMode",
		"CODING_MODE_5_1_4",
		&returns,
	)
	return returns
}

func Eac3AtmosCodingMode_CODING_MODE_7_1_4() Eac3AtmosCodingMode {
	_init_.Initialize()
	var returns Eac3AtmosCodingMode
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.Eac3AtmosCodingMode",
		"CODING_MODE_7_1_4",
		&returns,
	)
	return returns
}

func Eac3AtmosCodingMode_CODING_MODE_9_1_6() Eac3AtmosCodingMode {
	_init_.Initialize()
	var returns Eac3AtmosCodingMode
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.Eac3AtmosCodingMode",
		"CODING_MODE_9_1_6",
		&returns,
	)
	return returns
}

