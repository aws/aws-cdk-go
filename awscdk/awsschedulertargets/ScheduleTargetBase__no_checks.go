//go:build no_runtime_type_checking

package awsschedulertargets

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_ScheduleTargetBase) validateAddTargetActionToRoleParameters(role awsiam.IRole) error {
	return nil
}

func (s *jsiiProxy_ScheduleTargetBase) validateBindParameters(schedule awsscheduler.ISchedule) error {
	return nil
}

func (s *jsiiProxy_ScheduleTargetBase) validateBindBaseTargetConfigParameters(_schedule awsscheduler.ISchedule) error {
	return nil
}

func validateNewScheduleTargetBaseParameters(baseProps *ScheduleTargetBaseProps, targetArn *string) error {
	return nil
}

