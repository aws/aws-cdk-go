//go:build no_runtime_type_checking

package awscdkschedulertargetsalpha

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_ScheduleTargetBase) validateAddTargetActionToRoleParameters(role awsiam.IRole) error {
	return nil
}

func (s *jsiiProxy_ScheduleTargetBase) validateBindParameters(schedule awscdkscheduleralpha.ISchedule) error {
	return nil
}

func (s *jsiiProxy_ScheduleTargetBase) validateBindBaseTargetConfigParameters(_schedule awscdkscheduleralpha.ISchedule) error {
	return nil
}

func validateNewScheduleTargetBaseParameters(baseProps *ScheduleTargetBaseProps, targetArn *string) error {
	return nil
}

