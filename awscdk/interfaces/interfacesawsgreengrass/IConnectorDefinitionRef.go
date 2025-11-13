package interfacesawsgreengrass

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsgreengrass/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a ConnectorDefinition.
// Experimental.
type IConnectorDefinitionRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a ConnectorDefinition resource.
	// Experimental.
	ConnectorDefinitionRef() *ConnectorDefinitionReference
}

// The jsii proxy for IConnectorDefinitionRef
type jsiiProxy_IConnectorDefinitionRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (j *jsiiProxy_IConnectorDefinitionRef) ConnectorDefinitionRef() *ConnectorDefinitionReference {
	var returns *ConnectorDefinitionReference
	_jsii_.Get(
		j,
		"connectorDefinitionRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IConnectorDefinitionRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IConnectorDefinitionRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

