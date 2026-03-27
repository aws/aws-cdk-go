package interfacesawsroute53globalresolver

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsroute53globalresolver/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a GlobalResolver.
// Experimental.
type IGlobalResolverRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a GlobalResolver resource.
	// Experimental.
	GlobalResolverRef() *GlobalResolverReference
}

// The jsii proxy for IGlobalResolverRef
type jsiiProxy_IGlobalResolverRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (i *jsiiProxy_IGlobalResolverRef) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

func (j *jsiiProxy_IGlobalResolverRef) GlobalResolverRef() *GlobalResolverReference {
	var returns *GlobalResolverReference
	_jsii_.Get(
		j,
		"globalResolverRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IGlobalResolverRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IGlobalResolverRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

