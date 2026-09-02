package awsmedialivealpha

import (
	_init_ "github.com/aws/aws-cdk-go/awsmedialivealpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// EAC3 surround mode.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import medialive_alpha "github.com/aws/aws-cdk-go/awsmedialivealpha"
//
//   eac3SurroundMode := medialive_alpha.Eac3SurroundMode_DISABLED()
//
// Experimental.
type Eac3SurroundMode interface {
	// The underlying string value passed to CloudFormation.
	// Experimental.
	Value() *string
}

// The jsii proxy struct for Eac3SurroundMode
type jsiiProxy_Eac3SurroundMode struct {
	_ byte // padding
}

func (j *jsiiProxy_Eac3SurroundMode) Value() *string {
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
func Eac3SurroundMode_Of(value *string) Eac3SurroundMode {
	_init_.Initialize()

	if err := validateEac3SurroundMode_OfParameters(value); err != nil {
		panic(err)
	}
	var returns Eac3SurroundMode

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-medialive-alpha.Eac3SurroundMode",
		"of",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func Eac3SurroundMode_DISABLED() Eac3SurroundMode {
	_init_.Initialize()
	var returns Eac3SurroundMode
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.Eac3SurroundMode",
		"DISABLED",
		&returns,
	)
	return returns
}

func Eac3SurroundMode_ENABLED() Eac3SurroundMode {
	_init_.Initialize()
	var returns Eac3SurroundMode
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.Eac3SurroundMode",
		"ENABLED",
		&returns,
	)
	return returns
}

func Eac3SurroundMode_NOT_INDICATED() Eac3SurroundMode {
	_init_.Initialize()
	var returns Eac3SurroundMode
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.Eac3SurroundMode",
		"NOT_INDICATED",
		&returns,
	)
	return returns
}

