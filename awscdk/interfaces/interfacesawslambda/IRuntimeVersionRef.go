package interfacesawslambda

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawslambda/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a RuntimeVersion.
// Experimental.
type IRuntimeVersionRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a RuntimeVersion resource.
	// Experimental.
	RuntimeVersionRef() *RuntimeVersionReference
}

// The jsii proxy for IRuntimeVersionRef
type jsiiProxy_IRuntimeVersionRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (i *jsiiProxy_IRuntimeVersionRef) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

func (j *jsiiProxy_IRuntimeVersionRef) RuntimeVersionRef() *RuntimeVersionReference {
	var returns *RuntimeVersionReference
	_jsii_.Get(
		j,
		"runtimeVersionRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IRuntimeVersionRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IRuntimeVersionRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

