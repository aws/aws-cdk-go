//go:build no_runtime_type_checking

package awscodepipelineactions

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CodeDeployServerDeployAction) validateBindParameters(scope constructs.Construct, stage awscodepipeline.IStage, options *awscodepipeline.ActionBindOptions) error {
	return nil
}

func (c *jsiiProxy_CodeDeployServerDeployAction) validateBoundParameters(_scope constructs.Construct, _stage awscodepipeline.IStage, options *awscodepipeline.ActionBindOptions) error {
	return nil
}

func (c *jsiiProxy_CodeDeployServerDeployAction) validateOnStateChangeParameters(name *string, options *awsevents.RuleProps) error {
	return nil
}

func (c *jsiiProxy_CodeDeployServerDeployAction) validateVariableExpressionParameters(variableName *string) error {
	return nil
}

func validateNewCodeDeployServerDeployActionParameters(props *CodeDeployServerDeployActionProps) error {
	return nil
}

