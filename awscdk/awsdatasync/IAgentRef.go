package awsdatasync

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsdatasync/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Agent.
// Experimental.
type IAgentRef interface {
	constructs.IConstruct
}

// The jsii proxy for IAgentRef
type jsiiProxy_IAgentRef struct {
	internal.Type__constructsIConstruct
}

