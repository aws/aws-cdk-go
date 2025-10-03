package awsgreengrass

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsgreengrass/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a FunctionDefinition.
// Experimental.
type IFunctionDefinitionRef interface {
	constructs.IConstruct
}

// The jsii proxy for IFunctionDefinitionRef
type jsiiProxy_IFunctionDefinitionRef struct {
	internal.Type__constructsIConstruct
}

