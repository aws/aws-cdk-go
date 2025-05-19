//go:build no_runtime_type_checking

package awscodepipelineactions

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CodeStarConnectionsSourceAction) validateBindParameters(scope constructs.Construct, stage awscodepipeline.IStage, options *awscodepipeline.ActionBindOptions) error {
	return nil
}

func (c *jsiiProxy_CodeStarConnectionsSourceAction) validateBoundParameters(_scope constructs.Construct, _stage awscodepipeline.IStage, options *awscodepipeline.ActionBindOptions) error {
	return nil
}

func (c *jsiiProxy_CodeStarConnectionsSourceAction) validateOnStateChangeParameters(name *string, options *awsevents.RuleProps) error {
	return nil
}

func (c *jsiiProxy_CodeStarConnectionsSourceAction) validateVariableExpressionParameters(variableName *string) error {
	return nil
}

func validateNewCodeStarConnectionsSourceActionParameters(props *CodeStarConnectionsSourceActionProps) error {
	return nil
}

