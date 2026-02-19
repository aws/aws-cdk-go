package interfacesawswisdom

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawswisdom/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a AssistantAssociation.
// Experimental.
type IAssistantAssociationRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a AssistantAssociation resource.
	// Experimental.
	AssistantAssociationRef() *AssistantAssociationReference
}

// The jsii proxy for IAssistantAssociationRef
type jsiiProxy_IAssistantAssociationRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (i *jsiiProxy_IAssistantAssociationRef) With(mixins ...constructs.IMixin) constructs.IConstruct {
	args := []interface{}{}
	for _, a := range mixins {
		args = append(args, a)
	}

	var returns constructs.IConstruct

	_jsii_.Invoke(
		i,
		"with",
		args,
		&returns,
	)

	return returns
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

func (j *jsiiProxy_IAssistantAssociationRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IAssistantAssociationRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

