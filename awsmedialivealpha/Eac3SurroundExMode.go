package awsmedialivealpha

import (
	_init_ "github.com/aws/aws-cdk-go/awsmedialivealpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// EAC3 surround ex mode.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import medialive_alpha "github.com/aws/aws-cdk-go/awsmedialivealpha"
//
//   eac3SurroundExMode := medialive_alpha.Eac3SurroundExMode_DISABLED()
//
// Experimental.
type Eac3SurroundExMode interface {
	// The underlying string value passed to CloudFormation.
	// Experimental.
	Value() *string
}

// The jsii proxy struct for Eac3SurroundExMode
type jsiiProxy_Eac3SurroundExMode struct {
	_ byte // padding
}

func (j *jsiiProxy_Eac3SurroundExMode) Value() *string {
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
func Eac3SurroundExMode_Of(value *string) Eac3SurroundExMode {
	_init_.Initialize()

	if err := validateEac3SurroundExMode_OfParameters(value); err != nil {
		panic(err)
	}
	var returns Eac3SurroundExMode

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-medialive-alpha.Eac3SurroundExMode",
		"of",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func Eac3SurroundExMode_DISABLED() Eac3SurroundExMode {
	_init_.Initialize()
	var returns Eac3SurroundExMode
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.Eac3SurroundExMode",
		"DISABLED",
		&returns,
	)
	return returns
}

func Eac3SurroundExMode_ENABLED() Eac3SurroundExMode {
	_init_.Initialize()
	var returns Eac3SurroundExMode
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.Eac3SurroundExMode",
		"ENABLED",
		&returns,
	)
	return returns
}

func Eac3SurroundExMode_NOT_INDICATED() Eac3SurroundExMode {
	_init_.Initialize()
	var returns Eac3SurroundExMode
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.Eac3SurroundExMode",
		"NOT_INDICATED",
		&returns,
	)
	return returns
}

