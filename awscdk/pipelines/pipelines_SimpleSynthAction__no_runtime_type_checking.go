//go:build no_runtime_type_checking

package pipelines

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_SimpleSynthAction) validateBindParameters(scope awscdk.Construct, stage awscodepipeline.IStage, options *awscodepipeline.ActionBindOptions) error {
	return nil
}

func (s *jsiiProxy_SimpleSynthAction) validateOnStateChangeParameters(name *string, options *awsevents.RuleProps) error {
	return nil
}

func validateSimpleSynthAction_StandardNpmSynthParameters(options *StandardNpmSynthOptions) error {
	return nil
}

func validateSimpleSynthAction_StandardYarnSynthParameters(options *StandardYarnSynthOptions) error {
	return nil
}

func validateNewSimpleSynthActionParameters(props *SimpleSynthActionProps) error {
	return nil
}

