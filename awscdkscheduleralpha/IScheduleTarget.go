package awscdkscheduleralpha

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Interface representing a Event Bridge Schedule Target.
// Deprecated.
type IScheduleTarget interface {
	// Returns the schedule target specification.
	// Deprecated.
	Bind(_schedule ISchedule) *ScheduleTargetConfig
}

// The jsii proxy for IScheduleTarget
type jsiiProxy_IScheduleTarget struct {
	_ byte // padding
}

func (i *jsiiProxy_IScheduleTarget) Bind(_schedule ISchedule) *ScheduleTargetConfig {
	if err := i.validateBindParameters(_schedule); err != nil {
		panic(err)
	}
	var returns *ScheduleTargetConfig

	_jsii_.Invoke(
		i,
		"bind",
		[]interface{}{_schedule},
		&returns,
	)

	return returns
}

