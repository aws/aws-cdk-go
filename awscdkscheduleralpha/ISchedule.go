package awscdkscheduleralpha

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdkscheduleralpha/v2/internal"
)

// Interface representing a created or an imported `Schedule`.
// Experimental.
type ISchedule interface {
	awscdk.IResource
	// The schedule group associated with this schedule.
	// Experimental.
	Group() IGroup
	// The arn of the schedule.
	// Experimental.
	ScheduleArn() *string
	// The name of the schedule.
	// Experimental.
	ScheduleName() *string
}

// The jsii proxy for ISchedule
type jsiiProxy_ISchedule struct {
	internal.Type__awscdkIResource
}

func (j *jsiiProxy_ISchedule) Group() IGroup {
	var returns IGroup
	_jsii_.Get(
		j,
		"group",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ISchedule) ScheduleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"scheduleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ISchedule) ScheduleName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"scheduleName",
		&returns,
	)
	return returns
}

