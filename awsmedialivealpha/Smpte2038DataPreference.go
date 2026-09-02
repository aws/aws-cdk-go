package awsmedialivealpha

import (
	_init_ "github.com/aws/aws-cdk-go/awsmedialivealpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// SMPTE-2038 data preference.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import medialive_alpha "github.com/aws/aws-cdk-go/awsmedialivealpha"
//
//   smpte2038DataPreference := medialive_alpha.Smpte2038DataPreference_Of(jsii.String("value"))
//
// Experimental.
type Smpte2038DataPreference interface {
	// The underlying string value passed to CloudFormation.
	// Experimental.
	Value() *string
}

// The jsii proxy struct for Smpte2038DataPreference
type jsiiProxy_Smpte2038DataPreference struct {
	_ byte // padding
}

func (j *jsiiProxy_Smpte2038DataPreference) Value() *string {
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
func Smpte2038DataPreference_Of(value *string) Smpte2038DataPreference {
	_init_.Initialize()

	if err := validateSmpte2038DataPreference_OfParameters(value); err != nil {
		panic(err)
	}
	var returns Smpte2038DataPreference

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-medialive-alpha.Smpte2038DataPreference",
		"of",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func Smpte2038DataPreference_IGNORE() Smpte2038DataPreference {
	_init_.Initialize()
	var returns Smpte2038DataPreference
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.Smpte2038DataPreference",
		"IGNORE",
		&returns,
	)
	return returns
}

func Smpte2038DataPreference_PREFER() Smpte2038DataPreference {
	_init_.Initialize()
	var returns Smpte2038DataPreference
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.Smpte2038DataPreference",
		"PREFER",
		&returns,
	)
	return returns
}

