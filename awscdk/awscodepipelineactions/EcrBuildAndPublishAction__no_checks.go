//go:build no_runtime_type_checking

package awscodepipelineactions

// Building without runtime type checking enabled, so all the below just return nil

func (e *jsiiProxy_EcrBuildAndPublishAction) validateBindParameters(scope constructs.Construct, stage awscodepipeline.IStage, options *awscodepipeline.ActionBindOptions) error {
	return nil
}

func (e *jsiiProxy_EcrBuildAndPublishAction) validateBoundParameters(scope constructs.Construct, stage awscodepipeline.IStage, options *awscodepipeline.ActionBindOptions) error {
	return nil
}

func (e *jsiiProxy_EcrBuildAndPublishAction) validateOnStateChangeParameters(name *string, options *awsevents.RuleProps) error {
	return nil
}

func (e *jsiiProxy_EcrBuildAndPublishAction) validateVariableExpressionParameters(variableName *string) error {
	return nil
}

func validateNewEcrBuildAndPublishActionParameters(props *EcrBuildAndPublishActionProps) error {
	return nil
}

