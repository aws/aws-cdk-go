package awsservicecatalog

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsservicecatalog/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a StackSetConstraint.
// Experimental.
type IStackSetConstraintRef interface {
	constructs.IConstruct
}

// The jsii proxy for IStackSetConstraintRef
type jsiiProxy_IStackSetConstraintRef struct {
	internal.Type__constructsIConstruct
}

