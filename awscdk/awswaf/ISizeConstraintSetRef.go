package awswaf

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awswaf/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a SizeConstraintSet.
// Experimental.
type ISizeConstraintSetRef interface {
	constructs.IConstruct
}

// The jsii proxy for ISizeConstraintSetRef
type jsiiProxy_ISizeConstraintSetRef struct {
	internal.Type__constructsIConstruct
}

