package awsecs

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsecs/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a TaskDefinition.
// Experimental.
type ITaskDefinitionRef interface {
	constructs.IConstruct
}

// The jsii proxy for ITaskDefinitionRef
type jsiiProxy_ITaskDefinitionRef struct {
	internal.Type__constructsIConstruct
}

