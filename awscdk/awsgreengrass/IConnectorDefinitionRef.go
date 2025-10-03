package awsgreengrass

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsgreengrass/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a ConnectorDefinition.
// Experimental.
type IConnectorDefinitionRef interface {
	constructs.IConstruct
}

// The jsii proxy for IConnectorDefinitionRef
type jsiiProxy_IConnectorDefinitionRef struct {
	internal.Type__constructsIConstruct
}

