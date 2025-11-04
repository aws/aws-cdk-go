package awscdk

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a HookTypeConfig.
// Experimental.
type IHookTypeConfigRef interface {
	constructs.IConstruct
	IEnvironmentAware
	// A reference to a HookTypeConfig resource.
	// Experimental.
	HookTypeConfigRef() *HookTypeConfigReference
}

// The jsii proxy for IHookTypeConfigRef
type jsiiProxy_IHookTypeConfigRef struct {
	internal.Type__constructsIConstruct
	jsiiProxy_IEnvironmentAware
}

func (j *jsiiProxy_IHookTypeConfigRef) HookTypeConfigRef() *HookTypeConfigReference {
	var returns *HookTypeConfigReference
	_jsii_.Get(
		j,
		"hookTypeConfigRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IHookTypeConfigRef) Env() *ResourceEnvironment {
	var returns *ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IHookTypeConfigRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

