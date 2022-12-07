package awsapigateway

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

type IModel interface {
	// Returns the model name, such as 'myModel'.
	ModelId() *string
}

// The jsii proxy for IModel
type jsiiProxy_IModel struct {
	_ byte // padding
}

func (j *jsiiProxy_IModel) ModelId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"modelId",
		&returns,
	)
	return returns
}

