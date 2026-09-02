package awsmedialivealpha

import (
	_init_ "github.com/aws/aws-cdk-go/awsmedialivealpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Mode when quad SDI input is selected.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import medialive_alpha "github.com/aws/aws-cdk-go/awsmedialivealpha"
//
//   sdiMode := medialive_alpha.SdiMode_Of(jsii.String("value"))
//
// Experimental.
type SdiMode interface {
	// The underlying string value passed to CloudFormation.
	// Experimental.
	Value() *string
}

// The jsii proxy struct for SdiMode
type jsiiProxy_SdiMode struct {
	_ byte // padding
}

func (j *jsiiProxy_SdiMode) Value() *string {
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
func SdiMode_Of(value *string) SdiMode {
	_init_.Initialize()

	if err := validateSdiMode_OfParameters(value); err != nil {
		panic(err)
	}
	var returns SdiMode

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-medialive-alpha.SdiMode",
		"of",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func SdiMode_INTERLEAVE() SdiMode {
	_init_.Initialize()
	var returns SdiMode
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.SdiMode",
		"INTERLEAVE",
		&returns,
	)
	return returns
}

func SdiMode_QUADRANT() SdiMode {
	_init_.Initialize()
	var returns SdiMode
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.SdiMode",
		"QUADRANT",
		&returns,
	)
	return returns
}

