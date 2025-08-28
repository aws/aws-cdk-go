//go:build no_runtime_type_checking

package awscodepipelineactions

// Building without runtime type checking enabled, so all the below just return nil

func (a *jsiiProxy_AlexaSkillDeployAction) validateBindParameters(scope constructs.Construct, stage awscodepipeline.IStage, options *awscodepipeline.ActionBindOptions) error {
	return nil
}

func (a *jsiiProxy_AlexaSkillDeployAction) validateBoundParameters(_scope constructs.Construct, _stage awscodepipeline.IStage, _options *awscodepipeline.ActionBindOptions) error {
	return nil
}

func (a *jsiiProxy_AlexaSkillDeployAction) validateOnStateChangeParameters(name *string, options *awsevents.RuleProps) error {
	return nil
}

func (a *jsiiProxy_AlexaSkillDeployAction) validateVariableExpressionParameters(variableName *string) error {
	return nil
}

func validateNewAlexaSkillDeployActionParameters(props *AlexaSkillDeployActionProps) error {
	return nil
}

