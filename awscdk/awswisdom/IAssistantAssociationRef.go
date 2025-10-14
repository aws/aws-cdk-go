package awswisdom

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awswisdom/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a AssistantAssociation.
// Experimental.
type IAssistantAssociationRef interface {
	constructs.IConstruct
	// A reference to a AssistantAssociation resource.
	// Experimental.
	AssistantAssociationRef() *AssistantAssociationReference
}

// The jsii proxy for IAssistantAssociationRef
type jsiiProxy_IAssistantAssociationRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IAssistantAssociationRef) AssistantAssociationRef() *AssistantAssociationReference {
	var returns *AssistantAssociationReference
	_jsii_.Get(
		j,
		"assistantAssociationRef",
		&returns,
	)
	return returns
}

