package awsgreengrass

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsgreengrass/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a ConnectorDefinitionVersion.
// Experimental.
type IConnectorDefinitionVersionRef interface {
	constructs.IConstruct
	// A reference to a ConnectorDefinitionVersion resource.
	// Experimental.
	ConnectorDefinitionVersionRef() *ConnectorDefinitionVersionReference
}

// The jsii proxy for IConnectorDefinitionVersionRef
type jsiiProxy_IConnectorDefinitionVersionRef struct {
	internal.Type__constructsIConstruct
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

