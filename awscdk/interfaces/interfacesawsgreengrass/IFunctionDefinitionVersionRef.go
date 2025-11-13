package interfacesawsgreengrass

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsgreengrass/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a FunctionDefinitionVersion.
// Experimental.
type IFunctionDefinitionVersionRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a FunctionDefinitionVersion resource.
	// Experimental.
	FunctionDefinitionVersionRef() *FunctionDefinitionVersionReference
}

// The jsii proxy for IFunctionDefinitionVersionRef
type jsiiProxy_IFunctionDefinitionVersionRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (j *jsiiProxy_IFunctionDefinitionVersionRef) FunctionDefinitionVersionRef() *FunctionDefinitionVersionReference {
	var returns *FunctionDefinitionVersionReference
	_jsii_.Get(
		j,
		"functionDefinitionVersionRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IFunctionDefinitionVersionRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IFunctionDefinitionVersionRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

