package awsmedialivealpha

import (
	_init_ "github.com/aws/aws-cdk-go/awsmedialivealpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Whether to upconvert 608 captions to 708.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import medialive_alpha "github.com/aws/aws-cdk-go/awsmedialivealpha"
//
//   convert608To708 := medialive_alpha.Convert608To708_Of(jsii.String("value"))
//
// Experimental.
type Convert608To708 interface {
	// The underlying string value passed to CloudFormation.
	// Experimental.
	Value() *string
}

// The jsii proxy struct for Convert608To708
type jsiiProxy_Convert608To708 struct {
	_ byte // padding
}

func (j *jsiiProxy_Convert608To708) Value() *string {
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
func Convert608To708_Of(value *string) Convert608To708 {
	_init_.Initialize()

	if err := validateConvert608To708_OfParameters(value); err != nil {
		panic(err)
	}
	var returns Convert608To708

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-medialive-alpha.Convert608To708",
		"of",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func Convert608To708_DISABLED() Convert608To708 {
	_init_.Initialize()
	var returns Convert608To708
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.Convert608To708",
		"DISABLED",
		&returns,
	)
	return returns
}

func Convert608To708_UPCONVERT() Convert608To708 {
	_init_.Initialize()
	var returns Convert608To708
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.Convert608To708",
		"UPCONVERT",
		&returns,
	)
	return returns
}

