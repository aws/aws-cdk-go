package awsappflow

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsappflow/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a ConnectorProfile.
// Experimental.
type IConnectorProfileRef interface {
	constructs.IConstruct
	// A reference to a ConnectorProfile resource.
	// Experimental.
	ConnectorProfileRef() *ConnectorProfileReference
}

// The jsii proxy for IConnectorProfileRef
type jsiiProxy_IConnectorProfileRef struct {
	internal.Type__constructsIConstruct
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

