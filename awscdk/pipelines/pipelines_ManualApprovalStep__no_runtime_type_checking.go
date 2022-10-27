//go:build no_runtime_type_checking

package pipelines

// Building without runtime type checking enabled, so all the below just return nil

func (m *jsiiProxy_ManualApprovalStep) validateAddDependencyFileSetParameters(fs FileSet) error {
	return nil
}

func (m *jsiiProxy_ManualApprovalStep) validateAddStepDependencyParameters(step Step) error {
	return nil
}

func (m *jsiiProxy_ManualApprovalStep) validateConfigurePrimaryOutputParameters(fs FileSet) error {
	return nil
}

func (m *jsiiProxy_ManualApprovalStep) validateDiscoverReferencedOutputsParameters(structure interface{}) error {
	return nil
}

func validateManualApprovalStep_SequenceParameters(steps *[]Step) error {
	return nil
}

func validateNewManualApprovalStepParameters(id *string, props *ManualApprovalStepProps) error {
	return nil
}

