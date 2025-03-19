//go:build no_runtime_type_checking

package awscdkschedulertargetsalpha

// Building without runtime type checking enabled, so all the below just return nil

func (e *jsiiProxy_EventBridgePutEvents) validateAddTargetActionToRoleParameters(role awsiam.IRole) error {
	return nil
}

func (e *jsiiProxy_EventBridgePutEvents) validateBindParameters(schedule awscdkscheduleralpha.ISchedule) error {
	return nil
}

func (e *jsiiProxy_EventBridgePutEvents) validateBindBaseTargetConfigParameters(_schedule awscdkscheduleralpha.ISchedule) error {
	return nil
}

func validateNewEventBridgePutEventsParameters(entry *EventBridgePutEventsEntry, props *ScheduleTargetBaseProps) error {
	return nil
}

