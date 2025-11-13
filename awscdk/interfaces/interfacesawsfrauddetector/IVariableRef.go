package interfacesawsfrauddetector

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsfrauddetector/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Variable.
// Experimental.
type IVariableRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a Variable resource.
	// Experimental.
	VariableRef() *VariableReference
}

// The jsii proxy for IVariableRef
type jsiiProxy_IVariableRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (j *jsiiProxy_IVariableRef) VariableRef() *VariableReference {
	var returns *VariableReference
	_jsii_.Get(
		j,
		"variableRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IVariableRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IVariableRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

