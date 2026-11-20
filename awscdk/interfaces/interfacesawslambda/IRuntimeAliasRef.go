package interfacesawslambda

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawslambda/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a RuntimeAlias.
// Experimental.
type IRuntimeAliasRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a RuntimeAlias resource.
	// Experimental.
	RuntimeAliasRef() *RuntimeAliasReference
}

// The jsii proxy for IRuntimeAliasRef
type jsiiProxy_IRuntimeAliasRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (i *jsiiProxy_IRuntimeAliasRef) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

func (j *jsiiProxy_IRuntimeAliasRef) RuntimeAliasRef() *RuntimeAliasReference {
	var returns *RuntimeAliasReference
	_jsii_.Get(
		j,
		"runtimeAliasRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IRuntimeAliasRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IRuntimeAliasRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

