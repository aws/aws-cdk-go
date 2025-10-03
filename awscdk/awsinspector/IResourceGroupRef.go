package awsinspector

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsinspector/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a ResourceGroup.
// Experimental.
type IResourceGroupRef interface {
	constructs.IConstruct
}

// The jsii proxy for IResourceGroupRef
type jsiiProxy_IResourceGroupRef struct {
	internal.Type__constructsIConstruct
}

