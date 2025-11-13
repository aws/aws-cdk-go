package interfacesawsgreengrass

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsgreengrass/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a FunctionDefinition.
// Experimental.
type IFunctionDefinitionRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a FunctionDefinition resource.
	// Experimental.
	FunctionDefinitionRef() *FunctionDefinitionReference
}

// The jsii proxy for IFunctionDefinitionRef
type jsiiProxy_IFunctionDefinitionRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (j *jsiiProxy_IFunctionDefinitionRef) FunctionDefinitionRef() *FunctionDefinitionReference {
	var returns *FunctionDefinitionReference
	_jsii_.Get(
		j,
		"functionDefinitionRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IFunctionDefinitionRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IFunctionDefinitionRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

