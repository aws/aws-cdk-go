//go:build no_runtime_type_checking

package awscodepipelineactions

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CloudFormationDeployStackSetAction) validateBindParameters(scope constructs.Construct, stage awscodepipeline.IStage, options *awscodepipeline.ActionBindOptions) error {
	return nil
}

func (c *jsiiProxy_CloudFormationDeployStackSetAction) validateBoundParameters(scope constructs.Construct, _stage awscodepipeline.IStage, options *awscodepipeline.ActionBindOptions) error {
	return nil
}

func (c *jsiiProxy_CloudFormationDeployStackSetAction) validateOnStateChangeParameters(name *string, options *awsevents.RuleProps) error {
	return nil
}

func (c *jsiiProxy_CloudFormationDeployStackSetAction) validateVariableExpressionParameters(variableName *string) error {
	return nil
}

func validateNewCloudFormationDeployStackSetActionParameters(props *CloudFormationDeployStackSetActionProps) error {
	return nil
}

