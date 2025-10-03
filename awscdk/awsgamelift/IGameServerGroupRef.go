package awsgamelift

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsgamelift/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a GameServerGroup.
// Experimental.
type IGameServerGroupRef interface {
	constructs.IConstruct
}

// The jsii proxy for IGameServerGroupRef
type jsiiProxy_IGameServerGroupRef struct {
	internal.Type__constructsIConstruct
}

