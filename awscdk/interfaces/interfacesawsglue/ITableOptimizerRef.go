package interfacesawsglue

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsglue/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a TableOptimizer.
// Experimental.
type ITableOptimizerRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a TableOptimizer resource.
	// Experimental.
	TableOptimizerRef() *TableOptimizerReference
}

// The jsii proxy for ITableOptimizerRef
type jsiiProxy_ITableOptimizerRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
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

func (j *jsiiProxy_ITableOptimizerRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ITableOptimizerRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

