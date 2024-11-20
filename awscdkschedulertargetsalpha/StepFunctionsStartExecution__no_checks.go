//go:build no_runtime_type_checking

package awscdkschedulertargetsalpha

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_StepFunctionsStartExecution) validateAddTargetActionToRoleParameters(role awsiam.IRole) error {
	return nil
}

func (s *jsiiProxy_StepFunctionsStartExecution) validateBindParameters(schedule awscdkscheduleralpha.ISchedule) error {
	return nil
}

func (s *jsiiProxy_StepFunctionsStartExecution) validateBindBaseTargetConfigParameters(_schedule awscdkscheduleralpha.ISchedule) error {
	return nil
}

func validateNewStepFunctionsStartExecutionParameters(stateMachine awsstepfunctions.IStateMachine, props *ScheduleTargetBaseProps) error {
	return nil
}

