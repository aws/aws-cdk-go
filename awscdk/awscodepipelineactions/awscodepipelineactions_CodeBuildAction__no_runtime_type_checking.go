//go:build no_runtime_type_checking

package awscodepipelineactions

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CodeBuildAction) validateBindParameters(scope constructs.Construct, stage awscodepipeline.IStage, options *awscodepipeline.ActionBindOptions) error {
	return nil
}

func (c *jsiiProxy_CodeBuildAction) validateBoundParameters(scope constructs.Construct, _stage awscodepipeline.IStage, options *awscodepipeline.ActionBindOptions) error {
	return nil
}

func (c *jsiiProxy_CodeBuildAction) validateOnStateChangeParameters(name *string, options *awsevents.RuleProps) error {
	return nil
}

func (c *jsiiProxy_CodeBuildAction) validateVariableParameters(variableName *string) error {
	return nil
}

func (c *jsiiProxy_CodeBuildAction) validateVariableExpressionParameters(variableName *string) error {
	return nil
}

func validateNewCodeBuildActionParameters(props *CodeBuildActionProps) error {
	return nil
}

