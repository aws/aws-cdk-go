package awsmedialivealpha

import (
	_init_ "github.com/aws/aws-cdk-go/awsmedialivealpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// KLV data passthrough behavior.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import medialive_alpha "github.com/aws/aws-cdk-go/awsmedialivealpha"
//
//   m2tsKlv := medialive_alpha.M2tsKlv_Of(jsii.String("value"))
//
// Experimental.
type M2tsKlv interface {
	// The underlying string value passed to CloudFormation.
	// Experimental.
	Value() *string
}

// The jsii proxy struct for M2tsKlv
type jsiiProxy_M2tsKlv struct {
	_ byte // padding
}

func (j *jsiiProxy_M2tsKlv) Value() *string {
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
func M2tsKlv_Of(value *string) M2tsKlv {
	_init_.Initialize()

	if err := validateM2tsKlv_OfParameters(value); err != nil {
		panic(err)
	}
	var returns M2tsKlv

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-medialive-alpha.M2tsKlv",
		"of",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func M2tsKlv_NONE() M2tsKlv {
	_init_.Initialize()
	var returns M2tsKlv
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.M2tsKlv",
		"NONE",
		&returns,
	)
	return returns
}

func M2tsKlv_PASSTHROUGH() M2tsKlv {
	_init_.Initialize()
	var returns M2tsKlv
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.M2tsKlv",
		"PASSTHROUGH",
		&returns,
	)
	return returns
}

