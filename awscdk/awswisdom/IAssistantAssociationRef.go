package awswisdom

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awswisdom/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a AssistantAssociation.
// Experimental.
type IAssistantAssociationRef interface {
	constructs.IConstruct
}

// The jsii proxy for IAssistantAssociationRef
type jsiiProxy_IAssistantAssociationRef struct {
	internal.Type__constructsIConstruct
}

