package interfacesawsapigateway

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsapigateway/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Method.
// Experimental.
type IMethodRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a Method resource.
	// Experimental.
	MethodRef() *MethodReference
}

// The jsii proxy for IMethodRef
type jsiiProxy_IMethodRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (j *jsiiProxy_IMethodRef) MethodRef() *MethodReference {
	var returns *MethodReference
	_jsii_.Get(
		j,
		"methodRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IMethodRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IMethodRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

