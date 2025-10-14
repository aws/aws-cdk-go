package awsglue

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsglue/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a TableOptimizer.
// Experimental.
type ITableOptimizerRef interface {
	constructs.IConstruct
	// A reference to a TableOptimizer resource.
	// Experimental.
	TableOptimizerRef() *TableOptimizerReference
}

// The jsii proxy for ITableOptimizerRef
type jsiiProxy_ITableOptimizerRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_ITableOptimizerRef) TableOptimizerRef() *TableOptimizerReference {
	var returns *TableOptimizerReference
	_jsii_.Get(
		j,
		"tableOptimizerRef",
		&returns,
	)
	return returns
}

