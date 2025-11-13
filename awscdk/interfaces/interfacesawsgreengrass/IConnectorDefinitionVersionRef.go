package interfacesawsgreengrass

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsgreengrass/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a ConnectorDefinitionVersion.
// Experimental.
type IConnectorDefinitionVersionRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a ConnectorDefinitionVersion resource.
	// Experimental.
	ConnectorDefinitionVersionRef() *ConnectorDefinitionVersionReference
}

// The jsii proxy for IConnectorDefinitionVersionRef
type jsiiProxy_IConnectorDefinitionVersionRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (j *jsiiProxy_IConnectorDefinitionVersionRef) ConnectorDefinitionVersionRef() *ConnectorDefinitionVersionReference {
	var returns *ConnectorDefinitionVersionReference
	_jsii_.Get(
		j,
		"connectorDefinitionVersionRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IConnectorDefinitionVersionRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IConnectorDefinitionVersionRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

