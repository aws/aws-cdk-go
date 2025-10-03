package awscustomerprofiles

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awscustomerprofiles/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a CalculatedAttributeDefinition.
// Experimental.
type ICalculatedAttributeDefinitionRef interface {
	constructs.IConstruct
}

// The jsii proxy for ICalculatedAttributeDefinitionRef
type jsiiProxy_ICalculatedAttributeDefinitionRef struct {
	internal.Type__constructsIConstruct
}

