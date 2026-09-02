package awsmedialivealpha

import (
	_init_ "github.com/aws/aws-cdk-go/awsmedialivealpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// H.265 deblocking filter.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import medialive_alpha "github.com/aws/aws-cdk-go/awsmedialivealpha"
//
//   h265Deblocking := medialive_alpha.H265Deblocking_Of(jsii.String("value"))
//
// Experimental.
type H265Deblocking interface {
	// The underlying string value passed to CloudFormation.
	// Experimental.
	Value() *string
}

// The jsii proxy struct for H265Deblocking
type jsiiProxy_H265Deblocking struct {
	_ byte // padding
}

func (j *jsiiProxy_H265Deblocking) Value() *string {
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
func H265Deblocking_Of(value *string) H265Deblocking {
	_init_.Initialize()

	if err := validateH265Deblocking_OfParameters(value); err != nil {
		panic(err)
	}
	var returns H265Deblocking

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-medialive-alpha.H265Deblocking",
		"of",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func H265Deblocking_DISABLED() H265Deblocking {
	_init_.Initialize()
	var returns H265Deblocking
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.H265Deblocking",
		"DISABLED",
		&returns,
	)
	return returns
}

func H265Deblocking_ENABLED() H265Deblocking {
	_init_.Initialize()
	var returns H265Deblocking
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.H265Deblocking",
		"ENABLED",
		&returns,
	)
	return returns
}

