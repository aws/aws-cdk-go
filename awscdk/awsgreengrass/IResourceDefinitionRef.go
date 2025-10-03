package awsgreengrass

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsgreengrass/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a ResourceDefinition.
// Experimental.
type IResourceDefinitionRef interface {
	constructs.IConstruct
}

// The jsii proxy for IResourceDefinitionRef
type jsiiProxy_IResourceDefinitionRef struct {
	internal.Type__constructsIConstruct
}

