package awslocation

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awslocation/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a TrackerConsumer.
// Experimental.
type ITrackerConsumerRef interface {
	constructs.IConstruct
	// A reference to a TrackerConsumer resource.
	// Experimental.
	TrackerConsumerRef() *TrackerConsumerReference
}

// The jsii proxy for ITrackerConsumerRef
type jsiiProxy_ITrackerConsumerRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_ITrackerConsumerRef) TrackerConsumerRef() *TrackerConsumerReference {
	var returns *TrackerConsumerReference
	_jsii_.Get(
		j,
		"trackerConsumerRef",
		&returns,
	)
	return returns
}

