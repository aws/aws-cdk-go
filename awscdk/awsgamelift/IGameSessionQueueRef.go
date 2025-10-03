package awsgamelift

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsgamelift/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a GameSessionQueue.
// Experimental.
type IGameSessionQueueRef interface {
	constructs.IConstruct
}

// The jsii proxy for IGameSessionQueueRef
type jsiiProxy_IGameSessionQueueRef struct {
	internal.Type__constructsIConstruct
}

