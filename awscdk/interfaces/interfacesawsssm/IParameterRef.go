package interfacesawsssm

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsssm/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Parameter.
// Experimental.
type IParameterRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a Parameter resource.
	// Experimental.
	ParameterRef() *ParameterReference
}

// The jsii proxy for IParameterRef
type jsiiProxy_IParameterRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (i *jsiiProxy_IParameterRef) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

func (j *jsiiProxy_IParameterRef) ParameterRef() *ParameterReference {
	var returns *ParameterReference
	_jsii_.Get(
		j,
		"parameterRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IParameterRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IParameterRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

