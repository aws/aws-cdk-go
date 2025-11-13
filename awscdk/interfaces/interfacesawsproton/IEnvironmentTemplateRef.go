package interfacesawsproton

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsproton/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a EnvironmentTemplate.
// Experimental.
type IEnvironmentTemplateRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a EnvironmentTemplate resource.
	// Experimental.
	EnvironmentTemplateRef() *EnvironmentTemplateReference
}

// The jsii proxy for IEnvironmentTemplateRef
type jsiiProxy_IEnvironmentTemplateRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (j *jsiiProxy_IEnvironmentTemplateRef) EnvironmentTemplateRef() *EnvironmentTemplateReference {
	var returns *EnvironmentTemplateReference
	_jsii_.Get(
		j,
		"environmentTemplateRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IEnvironmentTemplateRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IEnvironmentTemplateRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

