//go:build no_runtime_type_checking

package awscodepipelineactions

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CodeDeployEcsDeployAction) validateBindParameters(scope constructs.Construct, stage awscodepipeline.IStage, options *awscodepipeline.ActionBindOptions) error {
	return nil
}

func (c *jsiiProxy_CodeDeployEcsDeployAction) validateBoundParameters(_scope constructs.Construct, _stage awscodepipeline.IStage, options *awscodepipeline.ActionBindOptions) error {
	return nil
}

func (c *jsiiProxy_CodeDeployEcsDeployAction) validateOnStateChangeParameters(name *string, options *awsevents.RuleProps) error {
	return nil
}

func (c *jsiiProxy_CodeDeployEcsDeployAction) validateVariableExpressionParameters(variableName *string) error {
	return nil
}

func validateNewCodeDeployEcsDeployActionParameters(props *CodeDeployEcsDeployActionProps) error {
	return nil
}

