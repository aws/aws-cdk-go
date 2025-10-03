package awsservicecatalog

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsservicecatalog/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a ResourceUpdateConstraint.
// Experimental.
type IResourceUpdateConstraintRef interface {
	constructs.IConstruct
}

// The jsii proxy for IResourceUpdateConstraintRef
type jsiiProxy_IResourceUpdateConstraintRef struct {
	internal.Type__constructsIConstruct
}

