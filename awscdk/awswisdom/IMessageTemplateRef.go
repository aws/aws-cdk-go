package awswisdom

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awswisdom/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a MessageTemplate.
// Experimental.
type IMessageTemplateRef interface {
	constructs.IConstruct
}

// The jsii proxy for IMessageTemplateRef
type jsiiProxy_IMessageTemplateRef struct {
	internal.Type__constructsIConstruct
}

