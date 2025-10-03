package awsiottwinmaker

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiottwinmaker/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a ComponentType.
// Experimental.
type IComponentTypeRef interface {
	constructs.IConstruct
}

// The jsii proxy for IComponentTypeRef
type jsiiProxy_IComponentTypeRef struct {
	internal.Type__constructsIConstruct
}

