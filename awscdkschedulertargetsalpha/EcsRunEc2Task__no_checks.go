//go:build no_runtime_type_checking

package awscdkschedulertargetsalpha

// Building without runtime type checking enabled, so all the below just return nil

func (e *jsiiProxy_EcsRunEc2Task) validateAddTargetActionToRoleParameters(role awsiam.IRole) error {
	return nil
}

func (e *jsiiProxy_EcsRunEc2Task) validateBindParameters(schedule awscdkscheduleralpha.ISchedule) error {
	return nil
}

func (e *jsiiProxy_EcsRunEc2Task) validateBindBaseTargetConfigParameters(_schedule awscdkscheduleralpha.ISchedule) error {
	return nil
}

func validateNewEcsRunEc2TaskParameters(cluster awsecs.ICluster, props *Ec2TaskProps) error {
	return nil
}

