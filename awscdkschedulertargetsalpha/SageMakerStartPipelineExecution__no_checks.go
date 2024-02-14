//go:build no_runtime_type_checking

package awscdkschedulertargetsalpha

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_SageMakerStartPipelineExecution) validateAddTargetActionToRoleParameters(schedule awscdkscheduleralpha.ISchedule, role awsiam.IRole) error {
	return nil
}

func (s *jsiiProxy_SageMakerStartPipelineExecution) validateBindParameters(schedule awscdkscheduleralpha.ISchedule) error {
	return nil
}

func (s *jsiiProxy_SageMakerStartPipelineExecution) validateBindBaseTargetConfigParameters(schedule awscdkscheduleralpha.ISchedule) error {
	return nil
}

func validateNewSageMakerStartPipelineExecutionParameters(pipeline awssagemaker.IPipeline, props *SageMakerStartPipelineExecutionProps) error {
	return nil
}

