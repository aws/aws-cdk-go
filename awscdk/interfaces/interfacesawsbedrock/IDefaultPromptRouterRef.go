package interfacesawsbedrock

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsbedrock/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a DefaultPromptRouter.
// Experimental.
type IDefaultPromptRouterRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a DefaultPromptRouter resource.
	// Experimental.
	DefaultPromptRouterRef() *DefaultPromptRouterReference
}

// The jsii proxy for IDefaultPromptRouterRef
type jsiiProxy_IDefaultPromptRouterRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (i *jsiiProxy_IDefaultPromptRouterRef) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

func (j *jsiiProxy_IDefaultPromptRouterRef) DefaultPromptRouterRef() *DefaultPromptRouterReference {
	var returns *DefaultPromptRouterReference
	_jsii_.Get(
		j,
		"defaultPromptRouterRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IDefaultPromptRouterRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IDefaultPromptRouterRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

