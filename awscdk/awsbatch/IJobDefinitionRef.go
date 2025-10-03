package awsbatch

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsbatch/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a JobDefinition.
// Experimental.
type IJobDefinitionRef interface {
	constructs.IConstruct
}

// The jsii proxy for IJobDefinitionRef
type jsiiProxy_IJobDefinitionRef struct {
	internal.Type__constructsIConstruct
}

