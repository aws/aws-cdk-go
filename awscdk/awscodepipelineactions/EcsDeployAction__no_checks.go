//go:build no_runtime_type_checking

package awscodepipelineactions

// Building without runtime type checking enabled, so all the below just return nil

func (e *jsiiProxy_EcsDeployAction) validateBindParameters(scope constructs.Construct, stage awscodepipeline.IStage, options *awscodepipeline.ActionBindOptions) error {
	return nil
}

func (e *jsiiProxy_EcsDeployAction) validateBoundParameters(_scope constructs.Construct, _stage awscodepipeline.IStage, options *awscodepipeline.ActionBindOptions) error {
	return nil
}

func (e *jsiiProxy_EcsDeployAction) validateOnStateChangeParameters(name *string, options *awsevents.RuleProps) error {
	return nil
}

func (e *jsiiProxy_EcsDeployAction) validateVariableExpressionParameters(variableName *string) error {
	return nil
}

func validateNewEcsDeployActionParameters(props *EcsDeployActionProps) error {
	return nil
}

