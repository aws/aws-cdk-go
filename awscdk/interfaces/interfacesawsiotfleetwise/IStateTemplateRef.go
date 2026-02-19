package interfacesawsiotfleetwise

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsiotfleetwise/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a StateTemplate.
// Experimental.
type IStateTemplateRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a StateTemplate resource.
	// Experimental.
	StateTemplateRef() *StateTemplateReference
}

// The jsii proxy for IStateTemplateRef
type jsiiProxy_IStateTemplateRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (i *jsiiProxy_IStateTemplateRef) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

func (j *jsiiProxy_IStateTemplateRef) StateTemplateRef() *StateTemplateReference {
	var returns *StateTemplateReference
	_jsii_.Get(
		j,
		"stateTemplateRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IStateTemplateRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IStateTemplateRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

