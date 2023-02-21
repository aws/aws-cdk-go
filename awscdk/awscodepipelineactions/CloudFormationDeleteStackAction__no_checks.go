//go:build no_runtime_type_checking

package awscodepipelineactions

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CloudFormationDeleteStackAction) validateAddToDeploymentRolePolicyParameters(statement awsiam.PolicyStatement) error {
	return nil
}

func (c *jsiiProxy_CloudFormationDeleteStackAction) validateBindParameters(scope constructs.Construct, stage awscodepipeline.IStage, options *awscodepipeline.ActionBindOptions) error {
	return nil
}

func (c *jsiiProxy_CloudFormationDeleteStackAction) validateBoundParameters(scope constructs.Construct, stage awscodepipeline.IStage, options *awscodepipeline.ActionBindOptions) error {
	return nil
}

func (c *jsiiProxy_CloudFormationDeleteStackAction) validateOnStateChangeParameters(name *string, options *awsevents.RuleProps) error {
	return nil
}

func (c *jsiiProxy_CloudFormationDeleteStackAction) validateVariableExpressionParameters(variableName *string) error {
	return nil
}

func validateNewCloudFormationDeleteStackActionParameters(props *CloudFormationDeleteStackActionProps) error {
	return nil
}

