package awswisdom

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awswisdom/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a MessageTemplate.
// Experimental.
type IMessageTemplateRef interface {
	constructs.IConstruct
	// A reference to a MessageTemplate resource.
	// Experimental.
	MessageTemplateRef() *MessageTemplateReference
}

// The jsii proxy for IMessageTemplateRef
type jsiiProxy_IMessageTemplateRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IMessageTemplateRef) MessageTemplateRef() *MessageTemplateReference {
	var returns *MessageTemplateReference
	_jsii_.Get(
		j,
		"messageTemplateRef",
		&returns,
	)
	return returns
}

