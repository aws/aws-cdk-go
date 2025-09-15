package awsivschat

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsivschat/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Room.
// Experimental.
type IRoomRef interface {
	constructs.IConstruct
	// A reference to a Room resource.
	// Experimental.
	RoomRef() *RoomReference
}

// The jsii proxy for IRoomRef
type jsiiProxy_IRoomRef struct {
	internal.Type__constructsIConstruct
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

