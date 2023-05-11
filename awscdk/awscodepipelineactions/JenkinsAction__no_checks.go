//go:build no_runtime_type_checking

package awscodepipelineactions

// Building without runtime type checking enabled, so all the below just return nil

func (j *jsiiProxy_JenkinsAction) validateBindParameters(scope constructs.Construct, stage awscodepipeline.IStage, options *awscodepipeline.ActionBindOptions) error {
	return nil
}

func (j *jsiiProxy_JenkinsAction) validateBoundParameters(_scope constructs.Construct, _stage awscodepipeline.IStage, _options *awscodepipeline.ActionBindOptions) error {
	return nil
}

func (j *jsiiProxy_JenkinsAction) validateOnStateChangeParameters(name *string, options *awsevents.RuleProps) error {
	return nil
}

func (j *jsiiProxy_JenkinsAction) validateVariableExpressionParameters(variableName *string) error {
	return nil
}

func validateNewJenkinsActionParameters(props *JenkinsActionProps) error {
	return nil
}

