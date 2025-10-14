package awswisdom

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awswisdom/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a MessageTemplateVersion.
// Experimental.
type IMessageTemplateVersionRef interface {
	constructs.IConstruct
	// A reference to a MessageTemplateVersion resource.
	// Experimental.
	MessageTemplateVersionRef() *MessageTemplateVersionReference
}

// The jsii proxy for IMessageTemplateVersionRef
type jsiiProxy_IMessageTemplateVersionRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IMessageTemplateVersionRef) MessageTemplateVersionRef() *MessageTemplateVersionReference {
	var returns *MessageTemplateVersionReference
	_jsii_.Get(
		j,
		"messageTemplateVersionRef",
		&returns,
	)
	return returns
}

