package awsconnect

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsconnect/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a AgentStatus.
// Experimental.
type IAgentStatusRef interface {
	constructs.IConstruct
}

// The jsii proxy for IAgentStatusRef
type jsiiProxy_IAgentStatusRef struct {
	internal.Type__constructsIConstruct
}

