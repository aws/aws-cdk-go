//go:build no_runtime_type_checking

package awsschedulertargets

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CodeBuildStartBuild) validateAddTargetActionToRoleParameters(role awsiam.IRole) error {
	return nil
}

func (c *jsiiProxy_CodeBuildStartBuild) validateBindParameters(schedule awsscheduler.ISchedule) error {
	return nil
}

func (c *jsiiProxy_CodeBuildStartBuild) validateBindBaseTargetConfigParameters(_schedule awsscheduler.ISchedule) error {
	return nil
}

func validateNewCodeBuildStartBuildParameters(project awscodebuild.IProject, props *ScheduleTargetBaseProps) error {
	return nil
}

