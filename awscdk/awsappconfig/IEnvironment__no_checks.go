//go:build no_runtime_type_checking

package awsappconfig

// Building without runtime type checking enabled, so all the below just return nil

func (i *jsiiProxy_IEnvironment) validateAddDeploymentParameters(configuration IConfiguration) error {
	return nil
}

func (i *jsiiProxy_IEnvironment) validateAddExtensionParameters(extension IExtension) error {
	return nil
}

func (i *jsiiProxy_IEnvironment) validateAtDeploymentTickParameters(eventDestination IEventDestination, options *ExtensionOptions) error {
	return nil
}

func (i *jsiiProxy_IEnvironment) validateGrantParameters(grantee awsiam.IGrantable) error {
	return nil
}

func (i *jsiiProxy_IEnvironment) validateGrantReadConfigParameters(grantee awsiam.IGrantable) error {
	return nil
}

func (i *jsiiProxy_IEnvironment) validateOnParameters(actionPoint ActionPoint, eventDestination IEventDestination, options *ExtensionOptions) error {
	return nil
}

func (i *jsiiProxy_IEnvironment) validateOnDeploymentBakingParameters(eventDestination IEventDestination, options *ExtensionOptions) error {
	return nil
}

func (i *jsiiProxy_IEnvironment) validateOnDeploymentCompleteParameters(eventDestination IEventDestination, options *ExtensionOptions) error {
	return nil
}

func (i *jsiiProxy_IEnvironment) validateOnDeploymentRolledBackParameters(eventDestination IEventDestination, options *ExtensionOptions) error {
	return nil
}

func (i *jsiiProxy_IEnvironment) validateOnDeploymentStartParameters(eventDestination IEventDestination, options *ExtensionOptions) error {
	return nil
}

func (i *jsiiProxy_IEnvironment) validateOnDeploymentStepParameters(eventDestination IEventDestination, options *ExtensionOptions) error {
	return nil
}

func (i *jsiiProxy_IEnvironment) validatePreCreateHostedConfigurationVersionParameters(eventDestination IEventDestination, options *ExtensionOptions) error {
	return nil
}

func (i *jsiiProxy_IEnvironment) validatePreStartDeploymentParameters(eventDestination IEventDestination, options *ExtensionOptions) error {
	return nil
}

func (i *jsiiProxy_IEnvironment) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

