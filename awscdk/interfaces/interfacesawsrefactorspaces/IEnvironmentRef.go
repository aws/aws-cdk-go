package interfacesawsrefactorspaces

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsrefactorspaces/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Environment.
// Experimental.
type IEnvironmentRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a Environment resource.
	// Experimental.
	EnvironmentRef() *EnvironmentReference
}

// The jsii proxy for IEnvironmentRef
type jsiiProxy_IEnvironmentRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (i *jsiiProxy_IEnvironmentRef) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

func (j *jsiiProxy_IEnvironmentRef) EnvironmentRef() *EnvironmentReference {
	var returns *EnvironmentReference
	_jsii_.Get(
		j,
		"environmentRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IEnvironmentRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IEnvironmentRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

