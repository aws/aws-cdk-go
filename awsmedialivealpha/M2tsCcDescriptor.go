package awsmedialivealpha

import (
	_init_ "github.com/aws/aws-cdk-go/awsmedialivealpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Whether to generate the captionServiceDescriptor in the PMT.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import medialive_alpha "github.com/aws/aws-cdk-go/awsmedialivealpha"
//
//   m2tsCcDescriptor := medialive_alpha.M2tsCcDescriptor_Of(jsii.String("value"))
//
// Experimental.
type M2tsCcDescriptor interface {
	// The underlying string value passed to CloudFormation.
	// Experimental.
	Value() *string
}

// The jsii proxy struct for M2tsCcDescriptor
type jsiiProxy_M2tsCcDescriptor struct {
	_ byte // padding
}

func (j *jsiiProxy_M2tsCcDescriptor) Value() *string {
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
func M2tsCcDescriptor_Of(value *string) M2tsCcDescriptor {
	_init_.Initialize()

	if err := validateM2tsCcDescriptor_OfParameters(value); err != nil {
		panic(err)
	}
	var returns M2tsCcDescriptor

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-medialive-alpha.M2tsCcDescriptor",
		"of",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func M2tsCcDescriptor_DISABLED() M2tsCcDescriptor {
	_init_.Initialize()
	var returns M2tsCcDescriptor
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.M2tsCcDescriptor",
		"DISABLED",
		&returns,
	)
	return returns
}

func M2tsCcDescriptor_ENABLED() M2tsCcDescriptor {
	_init_.Initialize()
	var returns M2tsCcDescriptor
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.M2tsCcDescriptor",
		"ENABLED",
		&returns,
	)
	return returns
}

