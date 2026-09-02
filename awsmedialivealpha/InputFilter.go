package awsmedialivealpha

import (
	_init_ "github.com/aws/aws-cdk-go/awsmedialivealpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Input filter mode.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import medialive_alpha "github.com/aws/aws-cdk-go/awsmedialivealpha"
//
//   inputFilter := medialive_alpha.InputFilter_AUTO()
//
// Experimental.
type InputFilter interface {
	// The underlying string value passed to CloudFormation.
	// Experimental.
	Value() *string
}

// The jsii proxy struct for InputFilter
type jsiiProxy_InputFilter struct {
	_ byte // padding
}

func (j *jsiiProxy_InputFilter) Value() *string {
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
func InputFilter_Of(value *string) InputFilter {
	_init_.Initialize()

	if err := validateInputFilter_OfParameters(value); err != nil {
		panic(err)
	}
	var returns InputFilter

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-medialive-alpha.InputFilter",
		"of",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func InputFilter_AUTO() InputFilter {
	_init_.Initialize()
	var returns InputFilter
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.InputFilter",
		"AUTO",
		&returns,
	)
	return returns
}

func InputFilter_DISABLED() InputFilter {
	_init_.Initialize()
	var returns InputFilter
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.InputFilter",
		"DISABLED",
		&returns,
	)
	return returns
}

func InputFilter_FORCED() InputFilter {
	_init_.Initialize()
	var returns InputFilter
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.InputFilter",
		"FORCED",
		&returns,
	)
	return returns
}

