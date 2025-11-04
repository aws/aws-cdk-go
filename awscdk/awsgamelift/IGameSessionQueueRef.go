package awsgamelift

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsgamelift/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a GameSessionQueue.
// Experimental.
type IGameSessionQueueRef interface {
	constructs.IConstruct
	awscdk.IEnvironmentAware
	// A reference to a GameSessionQueue resource.
	// Experimental.
	GameSessionQueueRef() *GameSessionQueueReference
}

// The jsii proxy for IGameSessionQueueRef
type jsiiProxy_IGameSessionQueueRef struct {
	internal.Type__constructsIConstruct
	internal.Type__awscdkIEnvironmentAware
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

func (j *jsiiProxy_IGameSessionQueueRef) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IGameSessionQueueRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

