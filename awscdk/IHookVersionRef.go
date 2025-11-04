package awscdk

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a HookVersion.
// Experimental.
type IHookVersionRef interface {
	constructs.IConstruct
	IEnvironmentAware
	// A reference to a HookVersion resource.
	// Experimental.
	HookVersionRef() *HookVersionReference
}

// The jsii proxy for IHookVersionRef
type jsiiProxy_IHookVersionRef struct {
	internal.Type__constructsIConstruct
	jsiiProxy_IEnvironmentAware
}

func (j *jsiiProxy_IHookVersionRef) HookVersionRef() *HookVersionReference {
	var returns *HookVersionReference
	_jsii_.Get(
		j,
		"hookVersionRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IHookVersionRef) Env() *ResourceEnvironment {
	var returns *ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IHookVersionRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

