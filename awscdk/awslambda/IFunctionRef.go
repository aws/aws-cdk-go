package awslambda

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Function.
// Experimental.
type IFunctionRef interface {
	constructs.IConstruct
	awscdk.IEnvironmentAware
	// A reference to a Function resource.
	// Experimental.
	FunctionRef() *FunctionReference
}

// The jsii proxy for IFunctionRef
type jsiiProxy_IFunctionRef struct {
	internal.Type__constructsIConstruct
	internal.Type__awscdkIEnvironmentAware
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

func (j *jsiiProxy_IFunctionRef) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IFunctionRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

