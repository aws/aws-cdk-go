package awsgamelift

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsgamelift/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a MatchmakingConfiguration.
// Experimental.
type IMatchmakingConfigurationRef interface {
	constructs.IConstruct
	awscdk.IEnvironmentAware
	// A reference to a MatchmakingConfiguration resource.
	// Experimental.
	MatchmakingConfigurationRef() *MatchmakingConfigurationReference
}

// The jsii proxy for IMatchmakingConfigurationRef
type jsiiProxy_IMatchmakingConfigurationRef struct {
	internal.Type__constructsIConstruct
	internal.Type__awscdkIEnvironmentAware
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

func (j *jsiiProxy_IMatchmakingConfigurationRef) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IMatchmakingConfigurationRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

