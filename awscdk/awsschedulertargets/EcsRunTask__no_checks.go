//go:build no_runtime_type_checking

package awsschedulertargets

// Building without runtime type checking enabled, so all the below just return nil

func (e *jsiiProxy_EcsRunTask) validateAddTargetActionToRoleParameters(role awsiam.IRole) error {
	return nil
}

func (e *jsiiProxy_EcsRunTask) validateBindParameters(schedule awsscheduler.ISchedule) error {
	return nil
}

func (e *jsiiProxy_EcsRunTask) validateBindBaseTargetConfigParameters(_schedule awsscheduler.ISchedule) error {
	return nil
}

func validateNewEcsRunTaskParameters(cluster awsecs.ICluster, props *EcsRunTaskBaseProps) error {
	return nil
}

