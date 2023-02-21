//go:build no_runtime_type_checking

package awscodepipelineactions

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CloudFormationExecuteChangeSetAction) validateBindParameters(scope constructs.Construct, stage awscodepipeline.IStage, options *awscodepipeline.ActionBindOptions) error {
	return nil
}

func (c *jsiiProxy_CloudFormationExecuteChangeSetAction) validateBoundParameters(scope constructs.Construct, stage awscodepipeline.IStage, options *awscodepipeline.ActionBindOptions) error {
	return nil
}

func (c *jsiiProxy_CloudFormationExecuteChangeSetAction) validateOnStateChangeParameters(name *string, options *awsevents.RuleProps) error {
	return nil
}

func (c *jsiiProxy_CloudFormationExecuteChangeSetAction) validateVariableExpressionParameters(variableName *string) error {
	return nil
}

func validateNewCloudFormationExecuteChangeSetActionParameters(props *CloudFormationExecuteChangeSetActionProps) error {
	return nil
}

