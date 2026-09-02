package awsmedialivealpha

import (
	_init_ "github.com/aws/aws-cdk-go/awsmedialivealpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Whether to include the ES Rate field in the PES header.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import medialive_alpha "github.com/aws/aws-cdk-go/awsmedialivealpha"
//
//   m2tsEsRateInPes := medialive_alpha.M2tsEsRateInPes_Of(jsii.String("value"))
//
// Experimental.
type M2tsEsRateInPes interface {
	// The underlying string value passed to CloudFormation.
	// Experimental.
	Value() *string
}

// The jsii proxy struct for M2tsEsRateInPes
type jsiiProxy_M2tsEsRateInPes struct {
	_ byte // padding
}

func (j *jsiiProxy_M2tsEsRateInPes) Value() *string {
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
func M2tsEsRateInPes_Of(value *string) M2tsEsRateInPes {
	_init_.Initialize()

	if err := validateM2tsEsRateInPes_OfParameters(value); err != nil {
		panic(err)
	}
	var returns M2tsEsRateInPes

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-medialive-alpha.M2tsEsRateInPes",
		"of",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func M2tsEsRateInPes_EXCLUDE() M2tsEsRateInPes {
	_init_.Initialize()
	var returns M2tsEsRateInPes
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.M2tsEsRateInPes",
		"EXCLUDE",
		&returns,
	)
	return returns
}

func M2tsEsRateInPes_INCLUDE() M2tsEsRateInPes {
	_init_.Initialize()
	var returns M2tsEsRateInPes
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.M2tsEsRateInPes",
		"INCLUDE",
		&returns,
	)
	return returns
}

