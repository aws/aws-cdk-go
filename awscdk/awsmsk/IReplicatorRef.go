package awsmsk

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsmsk/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Replicator.
// Experimental.
type IReplicatorRef interface {
	constructs.IConstruct
	// A reference to a Replicator resource.
	// Experimental.
	ReplicatorRef() *ReplicatorReference
}

// The jsii proxy for IReplicatorRef
type jsiiProxy_IReplicatorRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IReplicatorRef) ReplicatorRef() *ReplicatorReference {
	var returns *ReplicatorReference
	_jsii_.Get(
		j,
		"replicatorRef",
		&returns,
	)
	return returns
}

