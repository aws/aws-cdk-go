//go:build no_runtime_type_checking

package awscodepipelineactions

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CodeCommitSourceAction) validateBindParameters(scope constructs.Construct, stage awscodepipeline.IStage, options *awscodepipeline.ActionBindOptions) error {
	return nil
}

func (c *jsiiProxy_CodeCommitSourceAction) validateBoundParameters(_scope constructs.Construct, stage awscodepipeline.IStage, options *awscodepipeline.ActionBindOptions) error {
	return nil
}

func (c *jsiiProxy_CodeCommitSourceAction) validateOnStateChangeParameters(name *string, options *awsevents.RuleProps) error {
	return nil
}

func (c *jsiiProxy_CodeCommitSourceAction) validateVariableExpressionParameters(variableName *string) error {
	return nil
}

func validateNewCodeCommitSourceActionParameters(props *CodeCommitSourceActionProps) error {
	return nil
}

