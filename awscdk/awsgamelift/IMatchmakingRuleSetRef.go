package awsgamelift

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsgamelift/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a MatchmakingRuleSet.
// Experimental.
type IMatchmakingRuleSetRef interface {
	constructs.IConstruct
	// A reference to a MatchmakingRuleSet resource.
	// Experimental.
	MatchmakingRuleSetRef() *MatchmakingRuleSetReference
}

// The jsii proxy for IMatchmakingRuleSetRef
type jsiiProxy_IMatchmakingRuleSetRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IMatchmakingRuleSetRef) MatchmakingRuleSetRef() *MatchmakingRuleSetReference {
	var returns *MatchmakingRuleSetReference
	_jsii_.Get(
		j,
		"matchmakingRuleSetRef",
		&returns,
	)
	return returns
}

