package awsgreengrass

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsgreengrass/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a FunctionDefinitionVersion.
// Experimental.
type IFunctionDefinitionVersionRef interface {
	constructs.IConstruct
}

// The jsii proxy for IFunctionDefinitionVersionRef
type jsiiProxy_IFunctionDefinitionVersionRef struct {
	internal.Type__constructsIConstruct
}

