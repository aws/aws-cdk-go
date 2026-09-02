package awsmedialivealpha

import (
	_init_ "github.com/aws/aws-cdk-go/awsmedialivealpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// H.265 alternative transfer function.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import medialive_alpha "github.com/aws/aws-cdk-go/awsmedialivealpha"
//
//   h265AlternativeTransferFunction := medialive_alpha.H265AlternativeTransferFunction_Of(jsii.String("value"))
//
// Experimental.
type H265AlternativeTransferFunction interface {
	// The underlying string value passed to CloudFormation.
	// Experimental.
	Value() *string
}

// The jsii proxy struct for H265AlternativeTransferFunction
type jsiiProxy_H265AlternativeTransferFunction struct {
	_ byte // padding
}

func (j *jsiiProxy_H265AlternativeTransferFunction) Value() *string {
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
func H265AlternativeTransferFunction_Of(value *string) H265AlternativeTransferFunction {
	_init_.Initialize()

	if err := validateH265AlternativeTransferFunction_OfParameters(value); err != nil {
		panic(err)
	}
	var returns H265AlternativeTransferFunction

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-medialive-alpha.H265AlternativeTransferFunction",
		"of",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func H265AlternativeTransferFunction_INSERT() H265AlternativeTransferFunction {
	_init_.Initialize()
	var returns H265AlternativeTransferFunction
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.H265AlternativeTransferFunction",
		"INSERT",
		&returns,
	)
	return returns
}

func H265AlternativeTransferFunction_OMIT() H265AlternativeTransferFunction {
	_init_.Initialize()
	var returns H265AlternativeTransferFunction
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.H265AlternativeTransferFunction",
		"OMIT",
		&returns,
	)
	return returns
}

