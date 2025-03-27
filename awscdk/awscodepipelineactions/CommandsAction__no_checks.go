//go:build no_runtime_type_checking

package awscodepipelineactions

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CommandsAction) validateBindParameters(scope constructs.Construct, stage awscodepipeline.IStage, options *awscodepipeline.ActionBindOptions) error {
	return nil
}

func (c *jsiiProxy_CommandsAction) validateBoundParameters(scope constructs.Construct, stage awscodepipeline.IStage, options *awscodepipeline.ActionBindOptions) error {
	return nil
}

func (c *jsiiProxy_CommandsAction) validateOnStateChangeParameters(name *string, options *awsevents.RuleProps) error {
	return nil
}

func (c *jsiiProxy_CommandsAction) validateVariableParameters(variableName *string) error {
	return nil
}

func (c *jsiiProxy_CommandsAction) validateVariableExpressionParameters(variableName *string) error {
	return nil
}

func validateNewCommandsActionParameters(props *CommandsActionProps) error {
	return nil
}

