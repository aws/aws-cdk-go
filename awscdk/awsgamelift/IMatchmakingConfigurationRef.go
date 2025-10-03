package awsgamelift

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsgamelift/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a MatchmakingConfiguration.
// Experimental.
type IMatchmakingConfigurationRef interface {
	constructs.IConstruct
}

// The jsii proxy for IMatchmakingConfigurationRef
type jsiiProxy_IMatchmakingConfigurationRef struct {
	internal.Type__constructsIConstruct
}

