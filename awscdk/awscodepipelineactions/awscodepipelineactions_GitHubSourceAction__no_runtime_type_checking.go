//go:build no_runtime_type_checking

package awscodepipelineactions

// Building without runtime type checking enabled, so all the below just return nil

func (g *jsiiProxy_GitHubSourceAction) validateBindParameters(scope awscdk.Construct, stage awscodepipeline.IStage, options *awscodepipeline.ActionBindOptions) error {
	return nil
}

func (g *jsiiProxy_GitHubSourceAction) validateBoundParameters(scope awscdk.Construct, stage awscodepipeline.IStage, _options *awscodepipeline.ActionBindOptions) error {
	return nil
}

func (g *jsiiProxy_GitHubSourceAction) validateOnStateChangeParameters(name *string, options *awsevents.RuleProps) error {
	return nil
}

func (g *jsiiProxy_GitHubSourceAction) validateVariableExpressionParameters(variableName *string) error {
	return nil
}

func validateNewGitHubSourceActionParameters(props *GitHubSourceActionProps) error {
	return nil
}

