package interfacesawsconnect

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsconnect/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a ViewVersion.
// Experimental.
type IViewVersionRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a ViewVersion resource.
	// Experimental.
	ViewVersionRef() *ViewVersionReference
}

// The jsii proxy for IViewVersionRef
type jsiiProxy_IViewVersionRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (i *jsiiProxy_IViewVersionRef) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

func (j *jsiiProxy_IViewVersionRef) ViewVersionRef() *ViewVersionReference {
	var returns *ViewVersionReference
	_jsii_.Get(
		j,
		"viewVersionRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IViewVersionRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IViewVersionRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

