package interfacesawsappstream

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsappstream/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a AppBlockBuilder.
// Experimental.
type IAppBlockBuilderRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a AppBlockBuilder resource.
	// Experimental.
	AppBlockBuilderRef() *AppBlockBuilderReference
}

// The jsii proxy for IAppBlockBuilderRef
type jsiiProxy_IAppBlockBuilderRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (i *jsiiProxy_IAppBlockBuilderRef) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

func (j *jsiiProxy_IAppBlockBuilderRef) AppBlockBuilderRef() *AppBlockBuilderReference {
	var returns *AppBlockBuilderReference
	_jsii_.Get(
		j,
		"appBlockBuilderRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IAppBlockBuilderRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IAppBlockBuilderRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

