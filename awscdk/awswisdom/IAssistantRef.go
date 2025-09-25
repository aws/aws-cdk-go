package awswisdom

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awswisdom/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Assistant.
// Experimental.
type IAssistantRef interface {
	constructs.IConstruct
	// A reference to a Assistant resource.
	// Experimental.
	AssistantRef() *AssistantReference
}

// The jsii proxy for IAssistantRef
type jsiiProxy_IAssistantRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IAssistantRef) AssistantRef() *AssistantReference {
	var returns *AssistantReference
	_jsii_.Get(
		j,
		"assistantRef",
		&returns,
	)
	return returns
}

