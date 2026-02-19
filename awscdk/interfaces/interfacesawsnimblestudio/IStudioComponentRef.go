package interfacesawsnimblestudio

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsnimblestudio/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a StudioComponent.
// Experimental.
type IStudioComponentRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a StudioComponent resource.
	// Experimental.
	StudioComponentRef() *StudioComponentReference
}

// The jsii proxy for IStudioComponentRef
type jsiiProxy_IStudioComponentRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (i *jsiiProxy_IStudioComponentRef) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

func (j *jsiiProxy_IStudioComponentRef) StudioComponentRef() *StudioComponentReference {
	var returns *StudioComponentReference
	_jsii_.Get(
		j,
		"studioComponentRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IStudioComponentRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IStudioComponentRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

