package interfacesawscodebuild

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawscodebuild/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Sandbox.
// Experimental.
type ISandboxRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a Sandbox resource.
	// Experimental.
	SandboxRef() *SandboxReference
}

// The jsii proxy for ISandboxRef
type jsiiProxy_ISandboxRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (i *jsiiProxy_ISandboxRef) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

func (j *jsiiProxy_ISandboxRef) SandboxRef() *SandboxReference {
	var returns *SandboxReference
	_jsii_.Get(
		j,
		"sandboxRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ISandboxRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ISandboxRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

