package interfacesawsmedialive

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsmedialive/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Multiplexprogram.
// Experimental.
type IMultiplexprogramRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a Multiplexprogram resource.
	// Experimental.
	MultiplexprogramRef() *MultiplexprogramReference
}

// The jsii proxy for IMultiplexprogramRef
type jsiiProxy_IMultiplexprogramRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (i *jsiiProxy_IMultiplexprogramRef) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

func (j *jsiiProxy_IMultiplexprogramRef) MultiplexprogramRef() *MultiplexprogramReference {
	var returns *MultiplexprogramReference
	_jsii_.Get(
		j,
		"multiplexprogramRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IMultiplexprogramRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IMultiplexprogramRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

