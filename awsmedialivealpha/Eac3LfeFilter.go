package awsmedialivealpha

import (
	_init_ "github.com/aws/aws-cdk-go/awsmedialivealpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// EAC3 LFE filter.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import medialive_alpha "github.com/aws/aws-cdk-go/awsmedialivealpha"
//
//   eac3LfeFilter := medialive_alpha.Eac3LfeFilter_Of(jsii.String("value"))
//
// Experimental.
type Eac3LfeFilter interface {
	// The underlying string value passed to CloudFormation.
	// Experimental.
	Value() *string
}

// The jsii proxy struct for Eac3LfeFilter
type jsiiProxy_Eac3LfeFilter struct {
	_ byte // padding
}

func (j *jsiiProxy_Eac3LfeFilter) Value() *string {
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
func Eac3LfeFilter_Of(value *string) Eac3LfeFilter {
	_init_.Initialize()

	if err := validateEac3LfeFilter_OfParameters(value); err != nil {
		panic(err)
	}
	var returns Eac3LfeFilter

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-medialive-alpha.Eac3LfeFilter",
		"of",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func Eac3LfeFilter_DISABLED() Eac3LfeFilter {
	_init_.Initialize()
	var returns Eac3LfeFilter
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.Eac3LfeFilter",
		"DISABLED",
		&returns,
	)
	return returns
}

func Eac3LfeFilter_ENABLED() Eac3LfeFilter {
	_init_.Initialize()
	var returns Eac3LfeFilter
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.Eac3LfeFilter",
		"ENABLED",
		&returns,
	)
	return returns
}

