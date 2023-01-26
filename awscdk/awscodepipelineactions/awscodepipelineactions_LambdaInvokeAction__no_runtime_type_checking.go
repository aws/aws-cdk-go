//go:build no_runtime_type_checking

package awscodepipelineactions

// Building without runtime type checking enabled, so all the below just return nil

func (l *jsiiProxy_LambdaInvokeAction) validateBindParameters(scope constructs.Construct, stage awscodepipeline.IStage, options *awscodepipeline.ActionBindOptions) error {
	return nil
}

func (l *jsiiProxy_LambdaInvokeAction) validateBoundParameters(scope constructs.Construct, _stage awscodepipeline.IStage, options *awscodepipeline.ActionBindOptions) error {
	return nil
}

func (l *jsiiProxy_LambdaInvokeAction) validateOnStateChangeParameters(name *string, options *awsevents.RuleProps) error {
	return nil
}

func (l *jsiiProxy_LambdaInvokeAction) validateVariableParameters(variableName *string) error {
	return nil
}

func (l *jsiiProxy_LambdaInvokeAction) validateVariableExpressionParameters(variableName *string) error {
	return nil
}

func validateNewLambdaInvokeActionParameters(props *LambdaInvokeActionProps) error {
	return nil
}

