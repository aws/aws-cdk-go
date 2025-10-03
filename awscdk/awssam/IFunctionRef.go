package awssam

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awssam/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Function.
// Experimental.
type IFunctionRef interface {
	constructs.IConstruct
}

// The jsii proxy for IFunctionRef
type jsiiProxy_IFunctionRef struct {
	internal.Type__constructsIConstruct
}

