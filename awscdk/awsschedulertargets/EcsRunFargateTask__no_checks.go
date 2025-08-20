//go:build no_runtime_type_checking

package awsschedulertargets

// Building without runtime type checking enabled, so all the below just return nil

func (e *jsiiProxy_EcsRunFargateTask) validateAddTargetActionToRoleParameters(role awsiam.IRole) error {
	return nil
}

func (e *jsiiProxy_EcsRunFargateTask) validateBindParameters(schedule awsscheduler.ISchedule) error {
	return nil
}

func (e *jsiiProxy_EcsRunFargateTask) validateBindBaseTargetConfigParameters(_schedule awsscheduler.ISchedule) error {
	return nil
}

func validateNewEcsRunFargateTaskParameters(cluster awsecs.ICluster, props *FargateTaskProps) error {
	return nil
}

