//go:build no_runtime_type_checking

package pipelines

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CodeBuildStep) validateAddDependencyFileSetParameters(fs FileSet) error {
	return nil
}

func (c *jsiiProxy_CodeBuildStep) validateAddOutputDirectoryParameters(directory *string) error {
	return nil
}

func (c *jsiiProxy_CodeBuildStep) validateAddStepDependencyParameters(step Step) error {
	return nil
}

func (c *jsiiProxy_CodeBuildStep) validateConfigurePrimaryOutputParameters(fs FileSet) error {
	return nil
}

func (c *jsiiProxy_CodeBuildStep) validateDiscoverReferencedOutputsParameters(structure interface{}) error {
	return nil
}

func (c *jsiiProxy_CodeBuildStep) validateExportedVariableParameters(variableName *string) error {
	return nil
}

func (c *jsiiProxy_CodeBuildStep) validatePrimaryOutputDirectoryParameters(directory *string) error {
	return nil
}

func validateCodeBuildStep_SequenceParameters(steps *[]Step) error {
	return nil
}

func validateNewCodeBuildStepParameters(id *string, props *CodeBuildStepProps) error {
	return nil
}

