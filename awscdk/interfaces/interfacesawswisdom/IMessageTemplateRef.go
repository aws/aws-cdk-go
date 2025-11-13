package interfacesawswisdom

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawswisdom/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a MessageTemplate.
// Experimental.
type IMessageTemplateRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a MessageTemplate resource.
	// Experimental.
	MessageTemplateRef() *MessageTemplateReference
}

// The jsii proxy for IMessageTemplateRef
type jsiiProxy_IMessageTemplateRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
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

func (j *jsiiProxy_IMessageTemplateRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IMessageTemplateRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

