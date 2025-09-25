//go:build !no_runtime_type_checking

package awsschedulertargets

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsscheduler"
)

func (e *jsiiProxy_EventBridgePutEvents) validateAddTargetActionToRoleParameters(role awsiam.IRole) error {
	if role == nil {
		return fmt.Errorf("parameter role is required, but nil was provided")
	}

	return nil
}

func (e *jsiiProxy_EventBridgePutEvents) validateBindParameters(schedule awsscheduler.ISchedule) error {
	if schedule == nil {
		return fmt.Errorf("parameter schedule is required, but nil was provided")
	}

	return nil
}

func (e *jsiiProxy_EventBridgePutEvents) validateBindBaseTargetConfigParameters(_schedule awsscheduler.ISchedule) error {
	if _schedule == nil {
		return fmt.Errorf("parameter _schedule is required, but nil was provided")
	}

	return nil
}

func validateNewEventBridgePutEventsParameters(entry *EventBridgePutEventsEntry, props *ScheduleTargetBaseProps) error {
	if entry == nil {
		return fmt.Errorf("parameter entry is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(entry, func() string { return "parameter entry" }); err != nil {
		return err
	}

	if err := _jsii_.ValidateStruct(props, func() string { return "parameter props" }); err != nil {
		return err
	}

	return nil
}

