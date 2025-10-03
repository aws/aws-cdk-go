package awsgamelift

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsgamelift/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a MatchmakingRuleSet.
// Experimental.
type IMatchmakingRuleSetRef interface {
	constructs.IConstruct
}

// The jsii proxy for IMatchmakingRuleSetRef
type jsiiProxy_IMatchmakingRuleSetRef struct {
	internal.Type__constructsIConstruct
}

