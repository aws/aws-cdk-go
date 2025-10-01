package awslambda

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Function.
// Experimental.
type IFunctionRef interface {
	constructs.IConstruct
	// A reference to a Function resource.
	// Experimental.
	FunctionRef() *FunctionReference
}

// The jsii proxy for IFunctionRef
type jsiiProxy_IFunctionRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IFunctionRef) FunctionRef() *FunctionReference {
	var returns *FunctionReference
	_jsii_.Get(
		j,
		"functionRef",
		&returns,
	)
	return returns
}

