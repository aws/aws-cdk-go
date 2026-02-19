package interfacesawsathena

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsathena/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a WorkGroup.
// Experimental.
type IWorkGroupRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a WorkGroup resource.
	// Experimental.
	WorkGroupRef() *WorkGroupReference
}

// The jsii proxy for IWorkGroupRef
type jsiiProxy_IWorkGroupRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (i *jsiiProxy_IWorkGroupRef) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

func (j *jsiiProxy_IWorkGroupRef) WorkGroupRef() *WorkGroupReference {
	var returns *WorkGroupReference
	_jsii_.Get(
		j,
		"workGroupRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IWorkGroupRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IWorkGroupRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

