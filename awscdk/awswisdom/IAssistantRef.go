package awswisdom

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awswisdom/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Assistant.
// Experimental.
type IAssistantRef interface {
	constructs.IConstruct
}

// The jsii proxy for IAssistantRef
type jsiiProxy_IAssistantRef struct {
	internal.Type__constructsIConstruct
}

