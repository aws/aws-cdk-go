//go:build no_runtime_type_checking
// +build no_runtime_type_checking

package awscodepipelineactions

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_S3DeployAction) validateBindParameters(scope awscdk.Construct, stage awscodepipeline.IStage, options *awscodepipeline.ActionBindOptions) error {
	return nil
}

func (s *jsiiProxy_S3DeployAction) validateBoundParameters(_scope awscdk.Construct, _stage awscodepipeline.IStage, options *awscodepipeline.ActionBindOptions) error {
	return nil
}

func (s *jsiiProxy_S3DeployAction) validateOnStateChangeParameters(name *string, options *awsevents.RuleProps) error {
	return nil
}

func (s *jsiiProxy_S3DeployAction) validateVariableExpressionParameters(variableName *string) error {
	return nil
}

func validateNewS3DeployActionParameters(props *S3DeployActionProps) error {
	return nil
}

