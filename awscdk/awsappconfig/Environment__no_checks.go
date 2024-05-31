//go:build no_runtime_type_checking

package awsappconfig

// Building without runtime type checking enabled, so all the below just return nil

func (e *jsiiProxy_Environment) validateAddDeploymentParameters(configuration IConfiguration) error {
	return nil
}

func (e *jsiiProxy_Environment) validateAddExtensionParameters(extension IExtension) error {
	return nil
}

func (e *jsiiProxy_Environment) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (e *jsiiProxy_Environment) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (e *jsiiProxy_Environment) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func (e *jsiiProxy_Environment) validateGrantParameters(grantee awsiam.IGrantable) error {
	return nil
}

func (e *jsiiProxy_Environment) validateGrantReadConfigParameters(identity awsiam.IGrantable) error {
	return nil
}

func (e *jsiiProxy_Environment) validateOnParameters(actionPoint ActionPoint, eventDestination IEventDestination, options *ExtensionOptions) error {
	return nil
}

func (e *jsiiProxy_Environment) validateOnDeploymentBakingParameters(eventDestination IEventDestination, options *ExtensionOptions) error {
	return nil
}

func (e *jsiiProxy_Environment) validateOnDeploymentCompleteParameters(eventDestination IEventDestination, options *ExtensionOptions) error {
	return nil
}

func (e *jsiiProxy_Environment) validateOnDeploymentRolledBackParameters(eventDestination IEventDestination, options *ExtensionOptions) error {
	return nil
}

func (e *jsiiProxy_Environment) validateOnDeploymentStartParameters(eventDestination IEventDestination, options *ExtensionOptions) error {
	return nil
}

func (e *jsiiProxy_Environment) validateOnDeploymentStepParameters(eventDestination IEventDestination, options *ExtensionOptions) error {
	return nil
}

func (e *jsiiProxy_Environment) validatePreCreateHostedConfigurationVersionParameters(eventDestination IEventDestination, options *ExtensionOptions) error {
	return nil
}

func (e *jsiiProxy_Environment) validatePreStartDeploymentParameters(eventDestination IEventDestination, options *ExtensionOptions) error {
	return nil
}

func validateEnvironment_FromEnvironmentArnParameters(scope constructs.Construct, id *string, environmentArn *string) error {
	return nil
}

func validateEnvironment_FromEnvironmentAttributesParameters(scope constructs.Construct, id *string, attrs *EnvironmentAttributes) error {
	return nil
}

func validateEnvironment_IsConstructParameters(x interface{}) error {
	return nil
}

func validateEnvironment_IsOwnedResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateEnvironment_IsResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func (j *jsiiProxy_Environment) validateSetDeploymentQueueParameters(val *[]CfnDeployment) error {
	return nil
}

func (j *jsiiProxy_Environment) validateSetExtensibleParameters(val ExtensibleBase) error {
	return nil
}

func validateNewEnvironmentParameters(scope constructs.Construct, id *string, props *EnvironmentProps) error {
	return nil
}

