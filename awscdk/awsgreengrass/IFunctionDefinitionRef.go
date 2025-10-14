package awsgreengrass

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsgreengrass/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a FunctionDefinition.
// Experimental.
type IFunctionDefinitionRef interface {
	constructs.IConstruct
	// A reference to a FunctionDefinition resource.
	// Experimental.
	FunctionDefinitionRef() *FunctionDefinitionReference
}

// The jsii proxy for IFunctionDefinitionRef
type jsiiProxy_IFunctionDefinitionRef struct {
	internal.Type__constructsIConstruct
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

