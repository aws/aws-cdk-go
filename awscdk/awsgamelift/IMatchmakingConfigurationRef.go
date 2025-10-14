package awsgamelift

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsgamelift/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a MatchmakingConfiguration.
// Experimental.
type IMatchmakingConfigurationRef interface {
	constructs.IConstruct
	// A reference to a MatchmakingConfiguration resource.
	// Experimental.
	MatchmakingConfigurationRef() *MatchmakingConfigurationReference
}

// The jsii proxy for IMatchmakingConfigurationRef
type jsiiProxy_IMatchmakingConfigurationRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IMatchmakingConfigurationRef) MatchmakingConfigurationRef() *MatchmakingConfigurationReference {
	var returns *MatchmakingConfigurationReference
	_jsii_.Get(
		j,
		"matchmakingConfigurationRef",
		&returns,
	)
	return returns
}

