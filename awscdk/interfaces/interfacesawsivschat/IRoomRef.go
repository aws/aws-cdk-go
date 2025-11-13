package interfacesawsivschat

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsivschat/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Room.
// Experimental.
type IRoomRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a Room resource.
	// Experimental.
	RoomRef() *RoomReference
}

// The jsii proxy for IRoomRef
type jsiiProxy_IRoomRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (j *jsiiProxy_IRoomRef) RoomRef() *RoomReference {
	var returns *RoomReference
	_jsii_.Get(
		j,
		"roomRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IRoomRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IRoomRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

