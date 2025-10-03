package awsgreengrass

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsgreengrass/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a ResourceDefinitionVersion.
// Experimental.
type IResourceDefinitionVersionRef interface {
	constructs.IConstruct
}

// The jsii proxy for IResourceDefinitionVersionRef
type jsiiProxy_IResourceDefinitionVersionRef struct {
	internal.Type__constructsIConstruct
}

