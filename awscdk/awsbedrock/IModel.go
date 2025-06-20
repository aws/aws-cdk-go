package awsbedrock

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Represents a Bedrock model.
//
// The model could be a foundation model, a custom model, or a provisioned model.
type IModel interface {
	// The ARN of the model.
	// See: https://docs.aws.amazon.com/service-authorization/latest/reference/list_amazonbedrock.html#amazonbedrock-actions-as-permissions
	//
	ModelArn() *string
}

// The jsii proxy for IModel
type jsiiProxy_IModel struct {
	_ byte // padding
}

func (j *jsiiProxy_IModel) ModelArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"modelArn",
		&returns,
	)
	return returns
}

