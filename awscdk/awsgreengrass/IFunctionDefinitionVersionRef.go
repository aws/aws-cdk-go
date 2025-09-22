package awsgreengrass

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsgreengrass/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a FunctionDefinitionVersion.
// Experimental.
type IFunctionDefinitionVersionRef interface {
	constructs.IConstruct
	// A reference to a FunctionDefinitionVersion resource.
	// Experimental.
	FunctionDefinitionVersionRef() *FunctionDefinitionVersionReference
}

// The jsii proxy for IFunctionDefinitionVersionRef
type jsiiProxy_IFunctionDefinitionVersionRef struct {
	internal.Type__constructsIConstruct
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

