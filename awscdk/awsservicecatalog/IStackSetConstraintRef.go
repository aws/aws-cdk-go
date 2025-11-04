package awsservicecatalog

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsservicecatalog/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a StackSetConstraint.
// Experimental.
type IStackSetConstraintRef interface {
	constructs.IConstruct
	awscdk.IEnvironmentAware
	// A reference to a StackSetConstraint resource.
	// Experimental.
	StackSetConstraintRef() *StackSetConstraintReference
}

// The jsii proxy for IStackSetConstraintRef
type jsiiProxy_IStackSetConstraintRef struct {
	internal.Type__constructsIConstruct
	internal.Type__awscdkIEnvironmentAware
}

func (j *jsiiProxy_IStackSetConstraintRef) StackSetConstraintRef() *StackSetConstraintReference {
	var returns *StackSetConstraintReference
	_jsii_.Get(
		j,
		"stackSetConstraintRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IStackSetConstraintRef) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IStackSetConstraintRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

