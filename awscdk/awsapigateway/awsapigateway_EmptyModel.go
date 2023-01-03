package awsapigateway

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Represents a reference to a REST API's Empty model, which is available as part of the model collection by default.
//
// This can be used for mapping
// JSON responses from an integration to what is returned to a client,
// where strong typing is not required. In the absence of any defined
// model, the Empty model will be used to return the response payload
// unmapped.
//
// Definition
// {
//    "$schema" : "http://json-schema.org/draft-04/schema#",
//    "title" : "Empty Schema",
//    "type" : "object"
// }.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   emptyModel := awscdk.Aws_apigateway.NewEmptyModel()
//
// See: https://docs.amazonaws.cn/en_us/apigateway/latest/developerguide/models-mappings.html#models-mappings-models
//
// Deprecated: You should use Model.EMPTY_MODEL
type EmptyModel interface {
	IModel
	// Returns the model name, such as 'myModel'.
	// Deprecated: You should use Model.EMPTY_MODEL
	ModelId() *string
}

// The jsii proxy struct for EmptyModel
type jsiiProxy_EmptyModel struct {
	jsiiProxy_IModel
}

func (j *jsiiProxy_EmptyModel) ModelId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"modelId",
		&returns,
	)
	return returns
}


// Deprecated: You should use Model.EMPTY_MODEL
func NewEmptyModel() EmptyModel {
	_init_.Initialize()

	j := jsiiProxy_EmptyModel{}

	_jsii_.Create(
		"monocdk.aws_apigateway.EmptyModel",
		nil, // no parameters
		&j,
	)

	return &j
}

// Deprecated: You should use Model.EMPTY_MODEL
func NewEmptyModel_Override(e EmptyModel) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_apigateway.EmptyModel",
		nil, // no parameters
		e,
	)
}

