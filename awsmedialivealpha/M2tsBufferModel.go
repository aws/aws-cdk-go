package awsmedialivealpha

import (
	_init_ "github.com/aws/aws-cdk-go/awsmedialivealpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// The buffer model used for the transport stream.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import medialive_alpha "github.com/aws/aws-cdk-go/awsmedialivealpha"
//
//   m2tsBufferModel := medialive_alpha.M2tsBufferModel_Of(jsii.String("value"))
//
// Experimental.
type M2tsBufferModel interface {
	// The underlying string value passed to CloudFormation.
	// Experimental.
	Value() *string
}

// The jsii proxy struct for M2tsBufferModel
type jsiiProxy_M2tsBufferModel struct {
	_ byte // padding
}

func (j *jsiiProxy_M2tsBufferModel) Value() *string {
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
func M2tsBufferModel_Of(value *string) M2tsBufferModel {
	_init_.Initialize()

	if err := validateM2tsBufferModel_OfParameters(value); err != nil {
		panic(err)
	}
	var returns M2tsBufferModel

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-medialive-alpha.M2tsBufferModel",
		"of",
		[]interface{}{value},
		&returns,
	)

	return returns
}

func M2tsBufferModel_MULTIPLEX() M2tsBufferModel {
	_init_.Initialize()
	var returns M2tsBufferModel
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.M2tsBufferModel",
		"MULTIPLEX",
		&returns,
	)
	return returns
}

func M2tsBufferModel_NONE() M2tsBufferModel {
	_init_.Initialize()
	var returns M2tsBufferModel
	_jsii_.StaticGet(
		"@aws-cdk/aws-medialive-alpha.M2tsBufferModel",
		"NONE",
		&returns,
	)
	return returns
}

