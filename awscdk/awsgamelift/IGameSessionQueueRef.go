package awsgamelift

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsgamelift/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a GameSessionQueue.
// Experimental.
type IGameSessionQueueRef interface {
	constructs.IConstruct
	// A reference to a GameSessionQueue resource.
	// Experimental.
	GameSessionQueueRef() *GameSessionQueueReference
}

// The jsii proxy for IGameSessionQueueRef
type jsiiProxy_IGameSessionQueueRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IGameSessionQueueRef) GameSessionQueueRef() *GameSessionQueueReference {
	var returns *GameSessionQueueReference
	_jsii_.Get(
		j,
		"gameSessionQueueRef",
		&returns,
	)
	return returns
}

