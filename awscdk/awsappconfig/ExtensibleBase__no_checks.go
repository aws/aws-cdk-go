//go:build no_runtime_type_checking

package awsappconfig

// Building without runtime type checking enabled, so all the below just return nil

func (e *jsiiProxy_ExtensibleBase) validateAddExtensionParameters(extension IExtension) error {
	return nil
}

func (e *jsiiProxy_ExtensibleBase) validateAtDeploymentTickParameters(eventDestination IEventDestination, options *ExtensionOptions) error {
	return nil
}

func (e *jsiiProxy_ExtensibleBase) validateOnParameters(actionPoint ActionPoint, eventDestination IEventDestination, options *ExtensionOptions) error {
	return nil
}

func (e *jsiiProxy_ExtensibleBase) validateOnDeploymentBakingParameters(eventDestination IEventDestination, options *ExtensionOptions) error {
	return nil
}

func (e *jsiiProxy_ExtensibleBase) validateOnDeploymentCompleteParameters(eventDestination IEventDestination, options *ExtensionOptions) error {
	return nil
}

func (e *jsiiProxy_ExtensibleBase) validateOnDeploymentRolledBackParameters(eventDestination IEventDestination, options *ExtensionOptions) error {
	return nil
}

func (e *jsiiProxy_ExtensibleBase) validateOnDeploymentStartParameters(eventDestination IEventDestination, options *ExtensionOptions) error {
	return nil
}

func (e *jsiiProxy_ExtensibleBase) validateOnDeploymentStepParameters(eventDestination IEventDestination, options *ExtensionOptions) error {
	return nil
}

func (e *jsiiProxy_ExtensibleBase) validatePreCreateHostedConfigurationVersionParameters(eventDestination IEventDestination, options *ExtensionOptions) error {
	return nil
}

func (e *jsiiProxy_ExtensibleBase) validatePreStartDeploymentParameters(eventDestination IEventDestination, options *ExtensionOptions) error {
	return nil
}

func validateNewExtensibleBaseParameters(scope constructs.Construct, resourceArn *string) error {
	return nil
}

