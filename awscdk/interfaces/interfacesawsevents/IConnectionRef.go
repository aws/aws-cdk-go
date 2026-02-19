package interfacesawsevents

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsevents/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Connection.
// Experimental.
type IConnectionRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a Connection resource.
	// Experimental.
	ConnectionRef() *ConnectionReference
}

// The jsii proxy for IConnectionRef
type jsiiProxy_IConnectionRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (i *jsiiProxy_IConnectionRef) With(mixins ...constructs.IMixin) constructs.IConstruct {
	args := []interface{}{}
	for _, a := range mixins {
		args = append(args, a)
	}

	var returns constructs.IConstruct

	_jsii_.Invoke(
		i,
		"with",
		args,
		&returns,
	)

	return returns
}

func (j *jsiiProxy_IConnectionRef) ConnectionRef() *ConnectionReference {
	var returns *ConnectionReference
	_jsii_.Get(
		j,
		"connectionRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IConnectionRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IConnectionRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

