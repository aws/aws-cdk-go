package awsgamelift

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsgamelift/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a ContainerGroupDefinition.
// Experimental.
type IContainerGroupDefinitionRef interface {
	constructs.IConstruct
}

// The jsii proxy for IContainerGroupDefinitionRef
type jsiiProxy_IContainerGroupDefinitionRef struct {
	internal.Type__constructsIConstruct
}

