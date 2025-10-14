package awslocation

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awslocation/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Tracker.
// Experimental.
type ITrackerRef interface {
	constructs.IConstruct
	// A reference to a Tracker resource.
	// Experimental.
	TrackerRef() *TrackerReference
}

// The jsii proxy for ITrackerRef
type jsiiProxy_ITrackerRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_ITrackerRef) TrackerRef() *TrackerReference {
	var returns *TrackerReference
	_jsii_.Get(
		j,
		"trackerRef",
		&returns,
	)
	return returns
}

