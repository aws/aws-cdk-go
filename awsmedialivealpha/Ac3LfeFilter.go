package awsmedialivealpha

import (
	_init_ "github.com/aws/aws-cdk-go/awsmedialivealpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// AC3 LFE filter.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import medialive_alpha "github.com/aws/aws-cdk-go/awsmedialivealpha"
//
//   ac3LfeFilter := medialive_alpha.Ac3LfeFilter_Of(jsii.String("value"))
//
// Experimental.
type Ac3LfeFilter interface {
	// The underlying string value passed to CloudFormation.
	// Experimental.
	Value() *string
}

// The jsii proxy struct for Ac3LfeFilter
type jsiiProxy_Ac3LfeFilter struct {
	_ byte // padding
}

func (j *jsiiProxy_Ac3LfeFilter) Value() *string {
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
func Ac3LfeFilter_Of(value *string) Ac3LfeFilter {
	_init_.Initialize()

	if err := validateAc3LfeFilter_OfParameters(value); err != nil {
		panic(err)
	}
	var returns Ac3LfeFilter

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-medialive-alpha.Ac3LfeFilter",
		"of",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func Ac3LfeFilter_DISABLED() Ac3LfeFilter {
	_init_.Initialize()
	var returns Ac3LfeFilter
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.Ac3LfeFilter",
		"DISABLED",
		&returns,
	)
	return returns
}

func Ac3LfeFilter_ENABLED() Ac3LfeFilter {
	_init_.Initialize()
	var returns Ac3LfeFilter
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.Ac3LfeFilter",
		"ENABLED",
		&returns,
	)
	return returns
}

