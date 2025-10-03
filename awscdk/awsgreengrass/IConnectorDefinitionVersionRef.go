package awsgreengrass

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsgreengrass/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a ConnectorDefinitionVersion.
// Experimental.
type IConnectorDefinitionVersionRef interface {
	constructs.IConstruct
}

// The jsii proxy for IConnectorDefinitionVersionRef
type jsiiProxy_IConnectorDefinitionVersionRef struct {
	internal.Type__constructsIConstruct
}

