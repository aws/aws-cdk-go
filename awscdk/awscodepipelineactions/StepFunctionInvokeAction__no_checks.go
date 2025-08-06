//go:build no_runtime_type_checking

package awscodepipelineactions

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_StepFunctionInvokeAction) validateBindParameters(scope constructs.Construct, stage awscodepipeline.IStage, options *awscodepipeline.ActionBindOptions) error {
	return nil
}

func (s *jsiiProxy_StepFunctionInvokeAction) validateBoundParameters(_scope constructs.Construct, _stage awscodepipeline.IStage, options *awscodepipeline.ActionBindOptions) error {
	return nil
}

func (s *jsiiProxy_StepFunctionInvokeAction) validateOnStateChangeParameters(name *string, options *awsevents.RuleProps) error {
	return nil
}

func (s *jsiiProxy_StepFunctionInvokeAction) validateVariableExpressionParameters(variableName *string) error {
	return nil
}

func validateNewStepFunctionInvokeActionParameters(props *StepFunctionsInvokeActionProps) error {
	return nil
}

