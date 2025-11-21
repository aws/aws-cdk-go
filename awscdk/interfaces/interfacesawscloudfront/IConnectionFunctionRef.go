package interfacesawscloudfront

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawscloudfront/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a ConnectionFunction.
// Experimental.
type IConnectionFunctionRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a ConnectionFunction resource.
	// Experimental.
	ConnectionFunctionRef() *ConnectionFunctionReference
}

// The jsii proxy for IConnectionFunctionRef
type jsiiProxy_IConnectionFunctionRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (j *jsiiProxy_IConnectionFunctionRef) ConnectionFunctionRef() *ConnectionFunctionReference {
	var returns *ConnectionFunctionReference
	_jsii_.Get(
		j,
		"connectionFunctionRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IConnectionFunctionRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IConnectionFunctionRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

