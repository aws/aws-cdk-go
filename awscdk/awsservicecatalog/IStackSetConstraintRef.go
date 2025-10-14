package awsservicecatalog

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsservicecatalog/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a StackSetConstraint.
// Experimental.
type IStackSetConstraintRef interface {
	constructs.IConstruct
	// A reference to a StackSetConstraint resource.
	// Experimental.
	StackSetConstraintRef() *StackSetConstraintReference
}

// The jsii proxy for IStackSetConstraintRef
type jsiiProxy_IStackSetConstraintRef struct {
	internal.Type__constructsIConstruct
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

