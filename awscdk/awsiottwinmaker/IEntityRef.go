package awsiottwinmaker

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiottwinmaker/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Entity.
// Experimental.
type IEntityRef interface {
	constructs.IConstruct
}

// The jsii proxy for IEntityRef
type jsiiProxy_IEntityRef struct {
	internal.Type__constructsIConstruct
}

