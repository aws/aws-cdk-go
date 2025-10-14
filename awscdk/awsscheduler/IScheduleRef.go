package awsscheduler

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsscheduler/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Schedule.
// Experimental.
type IScheduleRef interface {
	constructs.IConstruct
	// A reference to a Schedule resource.
	// Experimental.
	ScheduleRef() *ScheduleReference
}

// The jsii proxy for IScheduleRef
type jsiiProxy_IScheduleRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IScheduleRef) ScheduleRef() *ScheduleReference {
	var returns *ScheduleReference
	_jsii_.Get(
		j,
		"scheduleRef",
		&returns,
	)
	return returns
}

