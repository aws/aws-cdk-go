//go:build no_runtime_type_checking

package awsschedulertargets

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_SageMakerStartPipelineExecution) validateAddTargetActionToRoleParameters(role awsiam.IRole) error {
	return nil
}

func (s *jsiiProxy_SageMakerStartPipelineExecution) validateBindParameters(schedule awsscheduler.ISchedule) error {
	return nil
}

func (s *jsiiProxy_SageMakerStartPipelineExecution) validateBindBaseTargetConfigParameters(_schedule awsscheduler.ISchedule) error {
	return nil
}

func validateNewSageMakerStartPipelineExecutionParameters(pipeline awssagemaker.IPipeline, props *SageMakerStartPipelineExecutionProps) error {
	return nil
}

