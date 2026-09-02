package awsmedialivealpha

import (
	_init_ "github.com/aws/aws-cdk-go/awsmedialivealpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// EBIF data passthrough behavior.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import medialive_alpha "github.com/aws/aws-cdk-go/awsmedialivealpha"
//
//   m2tsEbif := medialive_alpha.M2tsEbif_Of(jsii.String("value"))
//
// Experimental.
type M2tsEbif interface {
	// The underlying string value passed to CloudFormation.
	// Experimental.
	Value() *string
}

// The jsii proxy struct for M2tsEbif
type jsiiProxy_M2tsEbif struct {
	_ byte // padding
}

func (j *jsiiProxy_M2tsEbif) Value() *string {
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
func M2tsEbif_Of(value *string) M2tsEbif {
	_init_.Initialize()

	if err := validateM2tsEbif_OfParameters(value); err != nil {
		panic(err)
	}
	var returns M2tsEbif

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-medialive-alpha.M2tsEbif",
		"of",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func M2tsEbif_NONE() M2tsEbif {
	_init_.Initialize()
	var returns M2tsEbif
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.M2tsEbif",
		"NONE",
		&returns,
	)
	return returns
}

func M2tsEbif_PASSTHROUGH() M2tsEbif {
	_init_.Initialize()
	var returns M2tsEbif
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.M2tsEbif",
		"PASSTHROUGH",
		&returns,
	)
	return returns
}

