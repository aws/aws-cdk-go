package awstransfer

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awstransfer/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Connector.
// Experimental.
type IConnectorRef interface {
	constructs.IConstruct
	// A reference to a Connector resource.
	// Experimental.
	ConnectorRef() *ConnectorReference
}

// The jsii proxy for IConnectorRef
type jsiiProxy_IConnectorRef struct {
	internal.Type__constructsIConstruct
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

