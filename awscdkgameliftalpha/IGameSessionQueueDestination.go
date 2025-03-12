package awscdkgameliftalpha

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Represents a game session queue destination.
// Experimental.
type IGameSessionQueueDestination interface {
	// The ARN(s) to put into the destination field for a game session queue.
	//
	// This property is for cdk modules to consume only. You should not need to use this property.
	// Instead, use dedicated identifier on each components.
	// Experimental.
	ResourceArnForDestination() *string
}

// The jsii proxy for IGameSessionQueueDestination
type jsiiProxy_IGameSessionQueueDestination struct {
	_ byte // padding
}

func (j *jsiiProxy_IGameSessionQueueDestination) ResourceArnForDestination() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceArnForDestination",
		&returns,
	)
	return returns
}

