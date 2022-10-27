//go:build no_runtime_type_checking

package pipelines

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_ShellScriptAction) validateBindParameters(scope awscdk.Construct, stage awscodepipeline.IStage, options *awscodepipeline.ActionBindOptions) error {
	return nil
}

func (s *jsiiProxy_ShellScriptAction) validateOnStateChangeParameters(name *string, options *awsevents.RuleProps) error {
	return nil
}

func validateNewShellScriptActionParameters(props *ShellScriptActionProps) error {
	return nil
}

