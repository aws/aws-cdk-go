//go:build no_runtime_type_checking

package awsappconfig

// Building without runtime type checking enabled, so all the below just return nil

func (i *jsiiProxy_IApplication) validateAddEnvironmentParameters(id *string, options *EnvironmentOptions) error {
	return nil
}

func (i *jsiiProxy_IApplication) validateAddExistingEnvironmentParameters(environment IEnvironment) error {
	return nil
}

func (i *jsiiProxy_IApplication) validateAddExtensionParameters(extension IExtension) error {
	return nil
}

func (i *jsiiProxy_IApplication) validateAddHostedConfigurationParameters(id *string, options *HostedConfigurationOptions) error {
	return nil
}

func (i *jsiiProxy_IApplication) validateAddSourcedConfigurationParameters(id *string, options *SourcedConfigurationOptions) error {
	return nil
}

func (i *jsiiProxy_IApplication) validateAtDeploymentTickParameters(eventDestination IEventDestination, options *ExtensionOptions) error {
	return nil
}

func (i *jsiiProxy_IApplication) validateOnParameters(actionPoint ActionPoint, eventDestination IEventDestination, options *ExtensionOptions) error {
	return nil
}

func (i *jsiiProxy_IApplication) validateOnDeploymentBakingParameters(eventDestination IEventDestination, options *ExtensionOptions) error {
	return nil
}

func (i *jsiiProxy_IApplication) validateOnDeploymentCompleteParameters(eventDestination IEventDestination, options *ExtensionOptions) error {
	return nil
}

func (i *jsiiProxy_IApplication) validateOnDeploymentRolledBackParameters(eventDestination IEventDestination, options *ExtensionOptions) error {
	return nil
}

func (i *jsiiProxy_IApplication) validateOnDeploymentStartParameters(eventDestination IEventDestination, options *ExtensionOptions) error {
	return nil
}

func (i *jsiiProxy_IApplication) validateOnDeploymentStepParameters(eventDestination IEventDestination, options *ExtensionOptions) error {
	return nil
}

func (i *jsiiProxy_IApplication) validatePreCreateHostedConfigurationVersionParameters(eventDestination IEventDestination, options *ExtensionOptions) error {
	return nil
}

func (i *jsiiProxy_IApplication) validatePreStartDeploymentParameters(eventDestination IEventDestination, options *ExtensionOptions) error {
	return nil
}

