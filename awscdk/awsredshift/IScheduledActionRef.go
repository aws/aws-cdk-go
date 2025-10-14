package awsredshift

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsredshift/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a ScheduledAction.
// Experimental.
type IScheduledActionRef interface {
	constructs.IConstruct
	// A reference to a ScheduledAction resource.
	// Experimental.
	ScheduledActionRef() *ScheduledActionReference
}

// The jsii proxy for IScheduledActionRef
type jsiiProxy_IScheduledActionRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IScheduledActionRef) ScheduledActionRef() *ScheduledActionReference {
	var returns *ScheduledActionReference
	_jsii_.Get(
		j,
		"scheduledActionRef",
		&returns,
	)
	return returns
}

