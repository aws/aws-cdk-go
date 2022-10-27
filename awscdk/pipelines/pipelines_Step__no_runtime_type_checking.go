//go:build no_runtime_type_checking

package pipelines

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_Step) validateAddDependencyFileSetParameters(fs FileSet) error {
	return nil
}

func (s *jsiiProxy_Step) validateAddStepDependencyParameters(step Step) error {
	return nil
}

func (s *jsiiProxy_Step) validateConfigurePrimaryOutputParameters(fs FileSet) error {
	return nil
}

func (s *jsiiProxy_Step) validateDiscoverReferencedOutputsParameters(structure interface{}) error {
	return nil
}

func validateStep_SequenceParameters(steps *[]Step) error {
	return nil
}

func validateNewStepParameters(id *string) error {
	return nil
}

