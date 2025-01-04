//go:build no_runtime_type_checking

package awsappconfig

// Building without runtime type checking enabled, so all the below just return nil

func (h *jsiiProxy_HostedConfiguration) validateAddExtensionParameters(extension IExtension) error {
	return nil
}

func (h *jsiiProxy_HostedConfiguration) validateAtDeploymentTickParameters(eventDestination IEventDestination, options *ExtensionOptions) error {
	return nil
}

func (h *jsiiProxy_HostedConfiguration) validateOnParameters(actionPoint ActionPoint, eventDestination IEventDestination, options *ExtensionOptions) error {
	return nil
}

func (h *jsiiProxy_HostedConfiguration) validateOnDeploymentBakingParameters(eventDestination IEventDestination, options *ExtensionOptions) error {
	return nil
}

func (h *jsiiProxy_HostedConfiguration) validateOnDeploymentCompleteParameters(eventDestination IEventDestination, options *ExtensionOptions) error {
	return nil
}

func (h *jsiiProxy_HostedConfiguration) validateOnDeploymentRolledBackParameters(eventDestination IEventDestination, options *ExtensionOptions) error {
	return nil
}

func (h *jsiiProxy_HostedConfiguration) validateOnDeploymentStartParameters(eventDestination IEventDestination, options *ExtensionOptions) error {
	return nil
}

func (h *jsiiProxy_HostedConfiguration) validateOnDeploymentStepParameters(eventDestination IEventDestination, options *ExtensionOptions) error {
	return nil
}

func (h *jsiiProxy_HostedConfiguration) validatePreCreateHostedConfigurationVersionParameters(eventDestination IEventDestination, options *ExtensionOptions) error {
	return nil
}

func (h *jsiiProxy_HostedConfiguration) validatePreStartDeploymentParameters(eventDestination IEventDestination, options *ExtensionOptions) error {
	return nil
}

func validateHostedConfiguration_IsConstructParameters(x interface{}) error {
	return nil
}

func (j *jsiiProxy_HostedConfiguration) validateSetApplicationIdParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_HostedConfiguration) validateSetExtensibleParameters(val ExtensibleBase) error {
	return nil
}

func validateNewHostedConfigurationParameters(scope constructs.Construct, id *string, props *HostedConfigurationProps) error {
	return nil
}

