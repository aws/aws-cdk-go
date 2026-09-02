package awsmedialivealpha

import (
	_init_ "github.com/aws/aws-cdk-go/awsmedialivealpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// CBET insertion behavior when prior encoding is detected on the same layer.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import medialive_alpha "github.com/aws/aws-cdk-go/awsmedialivealpha"
//
//   nielsenCbetStepaside := medialive_alpha.NielsenCbetStepaside_Of(jsii.String("value"))
//
// Experimental.
type NielsenCbetStepaside interface {
	// The underlying string value passed to CloudFormation.
	// Experimental.
	Value() *string
}

// The jsii proxy struct for NielsenCbetStepaside
type jsiiProxy_NielsenCbetStepaside struct {
	_ byte // padding
}

func (j *jsiiProxy_NielsenCbetStepaside) Value() *string {
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
func NielsenCbetStepaside_Of(value *string) NielsenCbetStepaside {
	_init_.Initialize()

	if err := validateNielsenCbetStepaside_OfParameters(value); err != nil {
		panic(err)
	}
	var returns NielsenCbetStepaside

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-medialive-alpha.NielsenCbetStepaside",
		"of",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func NielsenCbetStepaside_DISABLED() NielsenCbetStepaside {
	_init_.Initialize()
	var returns NielsenCbetStepaside
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.NielsenCbetStepaside",
		"DISABLED",
		&returns,
	)
	return returns
}

func NielsenCbetStepaside_ENABLED() NielsenCbetStepaside {
	_init_.Initialize()
	var returns NielsenCbetStepaside
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.NielsenCbetStepaside",
		"ENABLED",
		&returns,
	)
	return returns
}

