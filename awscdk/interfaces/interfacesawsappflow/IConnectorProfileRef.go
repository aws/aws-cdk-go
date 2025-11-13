package interfacesawsappflow

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsappflow/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a ConnectorProfile.
// Experimental.
type IConnectorProfileRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a ConnectorProfile resource.
	// Experimental.
	ConnectorProfileRef() *ConnectorProfileReference
}

// The jsii proxy for IConnectorProfileRef
type jsiiProxy_IConnectorProfileRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (j *jsiiProxy_IConnectorProfileRef) ConnectorProfileRef() *ConnectorProfileReference {
	var returns *ConnectorProfileReference
	_jsii_.Get(
		j,
		"connectorProfileRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IConnectorProfileRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IConnectorProfileRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

