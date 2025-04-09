//go:build no_runtime_type_checking

package awsschedulertargets

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CodePipelineStartPipelineExecution) validateAddTargetActionToRoleParameters(role awsiam.IRole) error {
	return nil
}

func (c *jsiiProxy_CodePipelineStartPipelineExecution) validateBindParameters(schedule awsscheduler.ISchedule) error {
	return nil
}

func (c *jsiiProxy_CodePipelineStartPipelineExecution) validateBindBaseTargetConfigParameters(_schedule awsscheduler.ISchedule) error {
	return nil
}

func validateNewCodePipelineStartPipelineExecutionParameters(pipeline awscodepipeline.IPipeline, props *ScheduleTargetBaseProps) error {
	return nil
}

