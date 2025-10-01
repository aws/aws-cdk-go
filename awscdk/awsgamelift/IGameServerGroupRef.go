package awsgamelift

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsgamelift/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a GameServerGroup.
// Experimental.
type IGameServerGroupRef interface {
	constructs.IConstruct
	// A reference to a GameServerGroup resource.
	// Experimental.
	GameServerGroupRef() *GameServerGroupReference
}

// The jsii proxy for IGameServerGroupRef
type jsiiProxy_IGameServerGroupRef struct {
	internal.Type__constructsIConstruct
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

