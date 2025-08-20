package awsbedrockalpha

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Represents a Prompt Router, which provides intelligent routing between different models.
// Experimental.
type IPromptRouter interface {
	// The ARN of the prompt router.
	// Experimental.
	PromptRouterArn() *string
	// The ID of the prompt router.
	// Experimental.
	PromptRouterId() *string
	// The foundation models / profiles this router will route to.
	// Experimental.
	RoutingEndpoints() *[]IBedrockInvokable
}

// The jsii proxy for IPromptRouter
type jsiiProxy_IPromptRouter struct {
	_ byte // padding
}

func (j *jsiiProxy_IPromptRouter) PromptRouterArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"promptRouterArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IPromptRouter) PromptRouterId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"promptRouterId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IPromptRouter) RoutingEndpoints() *[]IBedrockInvokable {
	var returns *[]IBedrockInvokable
	_jsii_.Get(
		j,
		"routingEndpoints",
		&returns,
	)
	return returns
}

