//go:build no_runtime_type_checking

package appdelivery

// Building without runtime type checking enabled, so all the below just return nil

func (p *jsiiProxy_PipelineDeployStackAction) validateAddToDeploymentRolePolicyParameters(statement awsiam.PolicyStatement) error {
	return nil
}

func (p *jsiiProxy_PipelineDeployStackAction) validateBindParameters(scope awscdk.Construct, stage awscodepipeline.IStage, options *awscodepipeline.ActionBindOptions) error {
	return nil
}

func (p *jsiiProxy_PipelineDeployStackAction) validateOnStateChangeParameters(name *string, options *awsevents.RuleProps) error {
	return nil
}

func validateNewPipelineDeployStackActionParameters(props *PipelineDeployStackActionProps) error {
	return nil
}

