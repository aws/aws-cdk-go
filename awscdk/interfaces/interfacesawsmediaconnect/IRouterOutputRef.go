package interfacesawsmediaconnect

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsmediaconnect/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a RouterOutput.
// Experimental.
type IRouterOutputRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a RouterOutput resource.
	// Experimental.
	RouterOutputRef() *RouterOutputReference
}

// The jsii proxy for IRouterOutputRef
type jsiiProxy_IRouterOutputRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (i *jsiiProxy_IRouterOutputRef) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

func (j *jsiiProxy_IRouterOutputRef) RouterOutputRef() *RouterOutputReference {
	var returns *RouterOutputReference
	_jsii_.Get(
		j,
		"routerOutputRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IRouterOutputRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IRouterOutputRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

