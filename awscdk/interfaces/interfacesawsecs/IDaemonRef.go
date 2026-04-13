package interfacesawsecs

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsecs/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Daemon.
// Experimental.
type IDaemonRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a Daemon resource.
	// Experimental.
	DaemonRef() *DaemonReference
}

// The jsii proxy for IDaemonRef
type jsiiProxy_IDaemonRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (i *jsiiProxy_IDaemonRef) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

func (j *jsiiProxy_IDaemonRef) DaemonRef() *DaemonReference {
	var returns *DaemonReference
	_jsii_.Get(
		j,
		"daemonRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IDaemonRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IDaemonRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

