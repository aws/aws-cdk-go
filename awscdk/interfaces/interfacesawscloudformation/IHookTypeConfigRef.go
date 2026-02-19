package interfacesawscloudformation

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawscloudformation/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a HookTypeConfig.
// Experimental.
type IHookTypeConfigRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a HookTypeConfig resource.
	// Experimental.
	HookTypeConfigRef() *HookTypeConfigReference
}

// The jsii proxy for IHookTypeConfigRef
type jsiiProxy_IHookTypeConfigRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (i *jsiiProxy_IHookTypeConfigRef) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

func (j *jsiiProxy_IHookTypeConfigRef) HookTypeConfigRef() *HookTypeConfigReference {
	var returns *HookTypeConfigReference
	_jsii_.Get(
		j,
		"hookTypeConfigRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IHookTypeConfigRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
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

