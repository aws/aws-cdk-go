package awscdklocationalpha

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdklocationalpha/v2/internal"
)

// A Tracker.
// Experimental.
type ITracker interface {
	awscdk.IResource
	// The Amazon Resource Name (ARN) of the tracker resource.
	// Experimental.
	TrackerArn() *string
	// The name of the tracker.
	// Experimental.
	TrackerName() *string
}

// The jsii proxy for ITracker
type jsiiProxy_ITracker struct {
	internal.Type__awscdkIResource
}

func (j *jsiiProxy_ITracker) TrackerArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"trackerArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ITracker) TrackerName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"trackerName",
		&returns,
	)
	return returns
}

