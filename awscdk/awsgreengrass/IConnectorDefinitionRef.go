package awsgreengrass

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsgreengrass/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a ConnectorDefinition.
// Experimental.
type IConnectorDefinitionRef interface {
	constructs.IConstruct
	// A reference to a ConnectorDefinition resource.
	// Experimental.
	ConnectorDefinitionRef() *ConnectorDefinitionReference
}

// The jsii proxy for IConnectorDefinitionRef
type jsiiProxy_IConnectorDefinitionRef struct {
	internal.Type__constructsIConstruct
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

