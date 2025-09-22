//go:build no_runtime_type_checking

package awscodepipelineactions

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CloudFormationCreateReplaceChangeSetAction) validateAddToDeploymentRolePolicyParameters(statement awsiam.PolicyStatement) error {
	return nil
}

func (c *jsiiProxy_CloudFormationCreateReplaceChangeSetAction) validateBindParameters(scope constructs.Construct, stage awscodepipeline.IStage, options *awscodepipeline.ActionBindOptions) error {
	return nil
}

func (c *jsiiProxy_CloudFormationCreateReplaceChangeSetAction) validateBoundParameters(scope constructs.Construct, stage awscodepipeline.IStage, options *awscodepipeline.ActionBindOptions) error {
	return nil
}

func (c *jsiiProxy_CloudFormationCreateReplaceChangeSetAction) validateOnStateChangeParameters(name *string, options *awsevents.RuleProps) error {
	return nil
}

func (c *jsiiProxy_CloudFormationCreateReplaceChangeSetAction) validateVariableExpressionParameters(variableName *string) error {
	return nil
}

func validateNewCloudFormationCreateReplaceChangeSetActionParameters(props *CloudFormationCreateReplaceChangeSetActionProps) error {
	return nil
}

