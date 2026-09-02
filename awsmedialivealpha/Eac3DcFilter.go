package awsmedialivealpha

import (
	_init_ "github.com/aws/aws-cdk-go/awsmedialivealpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// EAC3 DC filter.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import medialive_alpha "github.com/aws/aws-cdk-go/awsmedialivealpha"
//
//   eac3DcFilter := medialive_alpha.Eac3DcFilter_Of(jsii.String("value"))
//
// Experimental.
type Eac3DcFilter interface {
	// The underlying string value passed to CloudFormation.
	// Experimental.
	Value() *string
}

// The jsii proxy struct for Eac3DcFilter
type jsiiProxy_Eac3DcFilter struct {
	_ byte // padding
}

func (j *jsiiProxy_Eac3DcFilter) Value() *string {
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
func Eac3DcFilter_Of(value *string) Eac3DcFilter {
	_init_.Initialize()

	if err := validateEac3DcFilter_OfParameters(value); err != nil {
		panic(err)
	}
	var returns Eac3DcFilter

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-medialive-alpha.Eac3DcFilter",
		"of",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func Eac3DcFilter_DISABLED() Eac3DcFilter {
	_init_.Initialize()
	var returns Eac3DcFilter
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.Eac3DcFilter",
		"DISABLED",
		&returns,
	)
	return returns
}

func Eac3DcFilter_ENABLED() Eac3DcFilter {
	_init_.Initialize()
	var returns Eac3DcFilter
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.Eac3DcFilter",
		"ENABLED",
		&returns,
	)
	return returns
}

