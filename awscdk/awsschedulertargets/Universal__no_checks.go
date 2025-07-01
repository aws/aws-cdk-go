//go:build no_runtime_type_checking

package awsschedulertargets

// Building without runtime type checking enabled, so all the below just return nil

func (u *jsiiProxy_Universal) validateAddTargetActionToRoleParameters(role awsiam.IRole) error {
	return nil
}

func (u *jsiiProxy_Universal) validateBindParameters(schedule awsscheduler.ISchedule) error {
	return nil
}

func (u *jsiiProxy_Universal) validateBindBaseTargetConfigParameters(_schedule awsscheduler.ISchedule) error {
	return nil
}

func validateNewUniversalParameters(props *UniversalTargetProps) error {
	return nil
}

