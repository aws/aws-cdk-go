//go:build no_runtime_type_checking

package awscodepipelineactions

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CloudFormationCreateUpdateStackAction) validateAddToDeploymentRolePolicyParameters(statement awsiam.PolicyStatement) error {
	return nil
}

func (c *jsiiProxy_CloudFormationCreateUpdateStackAction) validateBindParameters(scope constructs.Construct, stage awscodepipeline.IStage, options *awscodepipeline.ActionBindOptions) error {
	return nil
}

func (c *jsiiProxy_CloudFormationCreateUpdateStackAction) validateBoundParameters(scope constructs.Construct, stage awscodepipeline.IStage, options *awscodepipeline.ActionBindOptions) error {
	return nil
}

func (c *jsiiProxy_CloudFormationCreateUpdateStackAction) validateOnStateChangeParameters(name *string, options *awsevents.RuleProps) error {
	return nil
}

func (c *jsiiProxy_CloudFormationCreateUpdateStackAction) validateVariableExpressionParameters(variableName *string) error {
	return nil
}

func validateNewCloudFormationCreateUpdateStackActionParameters(props *CloudFormationCreateUpdateStackActionProps) error {
	return nil
}

