package awskafkaconnect

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awskafkaconnect/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Connector.
// Experimental.
type IConnectorRef interface {
	constructs.IConstruct
	awscdk.IEnvironmentAware
	// A reference to a Connector resource.
	// Experimental.
	ConnectorRef() *ConnectorReference
}

// The jsii proxy for IConnectorRef
type jsiiProxy_IConnectorRef struct {
	internal.Type__constructsIConstruct
	internal.Type__awscdkIEnvironmentAware
}

func (j *jsiiProxy_IConnectorRef) ConnectorRef() *ConnectorReference {
	var returns *ConnectorReference
	_jsii_.Get(
		j,
		"connectorRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IConnectorRef) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IConnectorRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

