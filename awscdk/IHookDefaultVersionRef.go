package awscdk

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a HookDefaultVersion.
// Experimental.
type IHookDefaultVersionRef interface {
	constructs.IConstruct
	IEnvironmentAware
	// A reference to a HookDefaultVersion resource.
	// Experimental.
	HookDefaultVersionRef() *HookDefaultVersionReference
}

// The jsii proxy for IHookDefaultVersionRef
type jsiiProxy_IHookDefaultVersionRef struct {
	internal.Type__constructsIConstruct
	jsiiProxy_IEnvironmentAware
}

func (j *jsiiProxy_IHookDefaultVersionRef) HookDefaultVersionRef() *HookDefaultVersionReference {
	var returns *HookDefaultVersionReference
	_jsii_.Get(
		j,
		"hookDefaultVersionRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IHookDefaultVersionRef) Env() *ResourceEnvironment {
	var returns *ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IHookDefaultVersionRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

