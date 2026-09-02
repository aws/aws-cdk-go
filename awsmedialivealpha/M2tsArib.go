package awsmedialivealpha

import (
	_init_ "github.com/aws/aws-cdk-go/awsmedialivealpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// ARIB-compliant field muxing.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import medialive_alpha "github.com/aws/aws-cdk-go/awsmedialivealpha"
//
//   m2tsArib := medialive_alpha.M2tsArib_Of(jsii.String("value"))
//
// Experimental.
type M2tsArib interface {
	// The underlying string value passed to CloudFormation.
	// Experimental.
	Value() *string
}

// The jsii proxy struct for M2tsArib
type jsiiProxy_M2tsArib struct {
	_ byte // padding
}

func (j *jsiiProxy_M2tsArib) Value() *string {
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
func M2tsArib_Of(value *string) M2tsArib {
	_init_.Initialize()

	if err := validateM2tsArib_OfParameters(value); err != nil {
		panic(err)
	}
	var returns M2tsArib

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-medialive-alpha.M2tsArib",
		"of",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func M2tsArib_DISABLED() M2tsArib {
	_init_.Initialize()
	var returns M2tsArib
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.M2tsArib",
		"DISABLED",
		&returns,
	)
	return returns
}

func M2tsArib_ENABLED() M2tsArib {
	_init_.Initialize()
	var returns M2tsArib
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.M2tsArib",
		"ENABLED",
		&returns,
	)
	return returns
}

