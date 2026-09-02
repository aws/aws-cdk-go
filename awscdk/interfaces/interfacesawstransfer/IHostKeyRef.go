package interfacesawstransfer

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawstransfer/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a HostKey.
// Experimental.
type IHostKeyRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a HostKey resource.
	// Experimental.
	HostKeyRef() *HostKeyReference
}

// The jsii proxy for IHostKeyRef
type jsiiProxy_IHostKeyRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (i *jsiiProxy_IHostKeyRef) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

func (j *jsiiProxy_IHostKeyRef) HostKeyRef() *HostKeyReference {
	var returns *HostKeyReference
	_jsii_.Get(
		j,
		"hostKeyRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IHostKeyRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IHostKeyRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

