//go:build no_runtime_type_checking

package awsschedulertargets

// Building without runtime type checking enabled, so all the below just return nil

func (l *jsiiProxy_LambdaInvoke) validateAddTargetActionToRoleParameters(role awsiam.IRole) error {
	return nil
}

func (l *jsiiProxy_LambdaInvoke) validateBindParameters(schedule awsscheduler.ISchedule) error {
	return nil
}

func (l *jsiiProxy_LambdaInvoke) validateBindBaseTargetConfigParameters(_schedule awsscheduler.ISchedule) error {
	return nil
}

func validateNewLambdaInvokeParameters(func_ awslambda.IFunction, props *ScheduleTargetBaseProps) error {
	return nil
}

