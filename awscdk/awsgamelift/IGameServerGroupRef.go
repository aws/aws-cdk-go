package awsgamelift

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsgamelift/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a GameServerGroup.
// Experimental.
type IGameServerGroupRef interface {
	constructs.IConstruct
	awscdk.IEnvironmentAware
	// A reference to a GameServerGroup resource.
	// Experimental.
	GameServerGroupRef() *GameServerGroupReference
}

// The jsii proxy for IGameServerGroupRef
type jsiiProxy_IGameServerGroupRef struct {
	internal.Type__constructsIConstruct
	internal.Type__awscdkIEnvironmentAware
}

func (j *jsiiProxy_IGameServerGroupRef) GameServerGroupRef() *GameServerGroupReference {
	var returns *GameServerGroupReference
	_jsii_.Get(
		j,
		"gameServerGroupRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IGameServerGroupRef) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IGameServerGroupRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

