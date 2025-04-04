//go:build no_runtime_type_checking

package awsappconfig

// Building without runtime type checking enabled, so all the below just return nil

func (i *jsiiProxy_IExtensible) validateAddExtensionParameters(extension IExtension) error {
	return nil
}

func (i *jsiiProxy_IExtensible) validateAtDeploymentTickParameters(eventDestination IEventDestination, options *ExtensionOptions) error {
	return nil
}

func (i *jsiiProxy_IExtensible) validateOnParameters(actionPoint ActionPoint, eventDestination IEventDestination, options *ExtensionOptions) error {
	return nil
}

func (i *jsiiProxy_IExtensible) validateOnDeploymentBakingParameters(eventDestination IEventDestination, options *ExtensionOptions) error {
	return nil
}

func (i *jsiiProxy_IExtensible) validateOnDeploymentCompleteParameters(eventDestination IEventDestination, options *ExtensionOptions) error {
	return nil
}

func (i *jsiiProxy_IExtensible) validateOnDeploymentRolledBackParameters(eventDestination IEventDestination, options *ExtensionOptions) error {
	return nil
}

func (i *jsiiProxy_IExtensible) validateOnDeploymentStartParameters(eventDestination IEventDestination, options *ExtensionOptions) error {
	return nil
}

func (i *jsiiProxy_IExtensible) validateOnDeploymentStepParameters(eventDestination IEventDestination, options *ExtensionOptions) error {
	return nil
}

func (i *jsiiProxy_IExtensible) validatePreCreateHostedConfigurationVersionParameters(eventDestination IEventDestination, options *ExtensionOptions) error {
	return nil
}

func (i *jsiiProxy_IExtensible) validatePreStartDeploymentParameters(eventDestination IEventDestination, options *ExtensionOptions) error {
	return nil
}

