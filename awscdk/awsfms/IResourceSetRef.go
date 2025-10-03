package awsfms

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsfms/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a ResourceSet.
// Experimental.
type IResourceSetRef interface {
	constructs.IConstruct
}

// The jsii proxy for IResourceSetRef
type jsiiProxy_IResourceSetRef struct {
	internal.Type__constructsIConstruct
}

