//go:build no_runtime_type_checking

package awscodepipelineactions

// Building without runtime type checking enabled, so all the below just return nil

func (a *jsiiProxy_Action) validateBindParameters(scope constructs.Construct, stage awscodepipeline.IStage, options *awscodepipeline.ActionBindOptions) error {
	return nil
}

func (a *jsiiProxy_Action) validateBoundParameters(scope constructs.Construct, stage awscodepipeline.IStage, options *awscodepipeline.ActionBindOptions) error {
	return nil
}

func (a *jsiiProxy_Action) validateOnStateChangeParameters(name *string, options *awsevents.RuleProps) error {
	return nil
}

func (a *jsiiProxy_Action) validateVariableExpressionParameters(variableName *string) error {
	return nil
}

func validateNewActionParameters(actionProperties *awscodepipeline.ActionProperties) error {
	return nil
}

