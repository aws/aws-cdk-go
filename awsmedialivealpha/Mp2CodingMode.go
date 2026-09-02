package awsmedialivealpha

import (
	_init_ "github.com/aws/aws-cdk-go/awsmedialivealpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// MP2 coding mode.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import medialive_alpha "github.com/aws/aws-cdk-go/awsmedialivealpha"
//
//   mp2CodingMode := medialive_alpha.Mp2CodingMode_Of(jsii.String("value"))
//
// Experimental.
type Mp2CodingMode interface {
	// The underlying string value passed to CloudFormation.
	// Experimental.
	Value() *string
}

// The jsii proxy struct for Mp2CodingMode
type jsiiProxy_Mp2CodingMode struct {
	_ byte // padding
}

func (j *jsiiProxy_Mp2CodingMode) Value() *string {
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
func Mp2CodingMode_Of(value *string) Mp2CodingMode {
	_init_.Initialize()

	if err := validateMp2CodingMode_OfParameters(value); err != nil {
		panic(err)
	}
	var returns Mp2CodingMode

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-medialive-alpha.Mp2CodingMode",
		"of",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func Mp2CodingMode_CODING_MODE_1_0() Mp2CodingMode {
	_init_.Initialize()
	var returns Mp2CodingMode
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.Mp2CodingMode",
		"CODING_MODE_1_0",
		&returns,
	)
	return returns
}

func Mp2CodingMode_CODING_MODE_2_0() Mp2CodingMode {
	_init_.Initialize()
	var returns Mp2CodingMode
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.Mp2CodingMode",
		"CODING_MODE_2_0",
		&returns,
	)
	return returns
}

