package awsbedrock

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsbedrock/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a AgentAlias.
// Experimental.
type IAgentAliasRef interface {
	constructs.IConstruct
}

// The jsii proxy for IAgentAliasRef
type jsiiProxy_IAgentAliasRef struct {
	internal.Type__constructsIConstruct
}

