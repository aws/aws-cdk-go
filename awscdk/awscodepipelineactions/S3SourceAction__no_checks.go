//go:build no_runtime_type_checking

package awscodepipelineactions

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_S3SourceAction) validateBindParameters(scope constructs.Construct, stage awscodepipeline.IStage, options *awscodepipeline.ActionBindOptions) error {
	return nil
}

func (s *jsiiProxy_S3SourceAction) validateBoundParameters(_scope constructs.Construct, stage awscodepipeline.IStage, options *awscodepipeline.ActionBindOptions) error {
	return nil
}

func (s *jsiiProxy_S3SourceAction) validateOnStateChangeParameters(name *string, options *awsevents.RuleProps) error {
	return nil
}

func (s *jsiiProxy_S3SourceAction) validateVariableExpressionParameters(variableName *string) error {
	return nil
}

func validateNewS3SourceActionParameters(props *S3SourceActionProps) error {
	return nil
}

