//go:build no_runtime_type_checking

package awscdkpipestargetsalpha

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_SageMakerTarget) validateBindParameters(pipe awscdkpipesalpha.IPipe) error {
	return nil
}

func (s *jsiiProxy_SageMakerTarget) validateGrantPushParameters(grantee awsiam.IRole) error {
	return nil
}

func validateNewSageMakerTargetParameters(pipeline awssagemaker.IPipeline, parameters *SageMakerTargetParameters) error {
	return nil
}

