//go:build no_runtime_type_checking

package awscodepipelineactions

// Building without runtime type checking enabled, so all the below just return nil

func (e *jsiiProxy_EcrSourceAction) validateBindParameters(scope constructs.Construct, stage awscodepipeline.IStage, options *awscodepipeline.ActionBindOptions) error {
	return nil
}

func (e *jsiiProxy_EcrSourceAction) validateBoundParameters(scope constructs.Construct, stage awscodepipeline.IStage, options *awscodepipeline.ActionBindOptions) error {
	return nil
}

func (e *jsiiProxy_EcrSourceAction) validateOnStateChangeParameters(name *string, options *awsevents.RuleProps) error {
	return nil
}

func (e *jsiiProxy_EcrSourceAction) validateVariableExpressionParameters(variableName *string) error {
	return nil
}

func validateNewEcrSourceActionParameters(props *EcrSourceActionProps) error {
	return nil
}

