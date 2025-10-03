package awsglue

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsglue/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a TableOptimizer.
// Experimental.
type ITableOptimizerRef interface {
	constructs.IConstruct
}

// The jsii proxy for ITableOptimizerRef
type jsiiProxy_ITableOptimizerRef struct {
	internal.Type__constructsIConstruct
}

