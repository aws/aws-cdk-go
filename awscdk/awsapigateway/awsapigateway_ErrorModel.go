package awsapigateway

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Represents a reference to a REST API's Error model, which is available as part of the model collection by default.
//
// This can be used for mapping
// error JSON responses from an integration to a client, where a simple
// generic message field is sufficient to map and return an error payload.
//
// Definition
// {
//    "$schema" : "http://json-schema.org/draft-04/schema#",
//    "title" : "Error Schema",
//    "type" : "object",
//    "properties" : {
//      "message" : { "type" : "string" }
//    }
// }.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   errorModel := awscdk.Aws_apigateway.NewErrorModel()
//
// Deprecated: You should use Model.ERROR_MODEL
type ErrorModel interface {
	IModel
	// Returns the model name, such as 'myModel'.
	// Deprecated: You should use Model.ERROR_MODEL
	ModelId() *string
}

// The jsii proxy struct for ErrorModel
type jsiiProxy_ErrorModel struct {
	jsiiProxy_IModel
}

func (j *jsiiProxy_ErrorModel) ModelId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"modelId",
		&returns,
	)
	return returns
}


// Deprecated: You should use Model.ERROR_MODEL
func NewErrorModel() ErrorModel {
	_init_.Initialize()

	j := jsiiProxy_ErrorModel{}

	_jsii_.Create(
		"monocdk.aws_apigateway.ErrorModel",
		nil, // no parameters
		&j,
	)

	return &j
}

// Deprecated: You should use Model.ERROR_MODEL
func NewErrorModel_Override(e ErrorModel) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_apigateway.ErrorModel",
		nil, // no parameters
		e,
	)
}

