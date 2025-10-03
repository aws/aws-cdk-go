package awscdk

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a ResourceVersion.
// Experimental.
type IResourceVersionRef interface {
	constructs.IConstruct
}

// The jsii proxy for IResourceVersionRef
type jsiiProxy_IResourceVersionRef struct {
	internal.Type__constructsIConstruct
}

