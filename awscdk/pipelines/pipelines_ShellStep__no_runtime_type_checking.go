//go:build no_runtime_type_checking

package pipelines

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_ShellStep) validateAddDependencyFileSetParameters(fs FileSet) error {
	return nil
}

func (s *jsiiProxy_ShellStep) validateAddOutputDirectoryParameters(directory *string) error {
	return nil
}

func (s *jsiiProxy_ShellStep) validateAddStepDependencyParameters(step Step) error {
	return nil
}

func (s *jsiiProxy_ShellStep) validateConfigurePrimaryOutputParameters(fs FileSet) error {
	return nil
}

func (s *jsiiProxy_ShellStep) validateDiscoverReferencedOutputsParameters(structure interface{}) error {
	return nil
}

func (s *jsiiProxy_ShellStep) validatePrimaryOutputDirectoryParameters(directory *string) error {
	return nil
}

func validateShellStep_SequenceParameters(steps *[]Step) error {
	return nil
}

func validateNewShellStepParameters(id *string, props *ShellStepProps) error {
	return nil
}

