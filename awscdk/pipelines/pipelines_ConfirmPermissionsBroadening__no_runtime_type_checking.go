//go:build no_runtime_type_checking

package pipelines

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_ConfirmPermissionsBroadening) validateAddDependencyFileSetParameters(fs FileSet) error {
	return nil
}

func (c *jsiiProxy_ConfirmPermissionsBroadening) validateAddStepDependencyParameters(step Step) error {
	return nil
}

func (c *jsiiProxy_ConfirmPermissionsBroadening) validateConfigurePrimaryOutputParameters(fs FileSet) error {
	return nil
}

func (c *jsiiProxy_ConfirmPermissionsBroadening) validateDiscoverReferencedOutputsParameters(structure interface{}) error {
	return nil
}

func (c *jsiiProxy_ConfirmPermissionsBroadening) validateProduceActionParameters(stage awscodepipeline.IStage, options *ProduceActionOptions) error {
	return nil
}

func validateConfirmPermissionsBroadening_SequenceParameters(steps *[]Step) error {
	return nil
}

func validateNewConfirmPermissionsBroadeningParameters(id *string, props *PermissionsBroadeningCheckProps) error {
	return nil
}

