//go:build no_runtime_type_checking

package awscodepipelineactions

// Building without runtime type checking enabled, so all the below just return nil

func (e *jsiiProxy_Ec2DeployAction) validateBindParameters(scope constructs.Construct, stage awscodepipeline.IStage, options *awscodepipeline.ActionBindOptions) error {
	return nil
}

func (e *jsiiProxy_Ec2DeployAction) validateBoundParameters(scope constructs.Construct, stage awscodepipeline.IStage, options *awscodepipeline.ActionBindOptions) error {
	return nil
}

func (e *jsiiProxy_Ec2DeployAction) validateOnStateChangeParameters(name *string, options *awsevents.RuleProps) error {
	return nil
}

func (e *jsiiProxy_Ec2DeployAction) validateVariableExpressionParameters(variableName *string) error {
	return nil
}

func validateNewEc2DeployActionParameters(props *Ec2DeployActionProps) error {
	return nil
}

