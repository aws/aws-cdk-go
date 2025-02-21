//go:build no_runtime_type_checking

package awscdkschedulertargetsalpha

// Building without runtime type checking enabled, so all the below just return nil

func (l *jsiiProxy_LambdaInvoke) validateAddTargetActionToRoleParameters(role awsiam.IRole) error {
	return nil
}

func (l *jsiiProxy_LambdaInvoke) validateBindParameters(schedule awscdkscheduleralpha.ISchedule) error {
	return nil
}

func (l *jsiiProxy_LambdaInvoke) validateBindBaseTargetConfigParameters(_schedule awscdkscheduleralpha.ISchedule) error {
	return nil
}

func validateNewLambdaInvokeParameters(func_ awslambda.IFunction, props *ScheduleTargetBaseProps) error {
	return nil
}

