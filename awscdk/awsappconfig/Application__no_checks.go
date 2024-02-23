//go:build no_runtime_type_checking

package awsappconfig

// Building without runtime type checking enabled, so all the below just return nil

func (a *jsiiProxy_Application) validateAddEnvironmentParameters(id *string, options *EnvironmentOptions) error {
	return nil
}

func (a *jsiiProxy_Application) validateAddExistingEnvironmentParameters(environment IEnvironment) error {
	return nil
}

func (a *jsiiProxy_Application) validateAddExtensionParameters(extension IExtension) error {
	return nil
}

func (a *jsiiProxy_Application) validateAddHostedConfigurationParameters(id *string, options *HostedConfigurationOptions) error {
	return nil
}

func (a *jsiiProxy_Application) validateAddSourcedConfigurationParameters(id *string, options *SourcedConfigurationOptions) error {
	return nil
}

func (a *jsiiProxy_Application) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (a *jsiiProxy_Application) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (a *jsiiProxy_Application) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func (a *jsiiProxy_Application) validateOnParameters(actionPoint ActionPoint, eventDestination IEventDestination, options *ExtensionOptions) error {
	return nil
}

func (a *jsiiProxy_Application) validateOnDeploymentBakingParameters(eventDestination IEventDestination, options *ExtensionOptions) error {
	return nil
}

func (a *jsiiProxy_Application) validateOnDeploymentCompleteParameters(eventDestination IEventDestination, options *ExtensionOptions) error {
	return nil
}

func (a *jsiiProxy_Application) validateOnDeploymentRolledBackParameters(eventDestination IEventDestination, options *ExtensionOptions) error {
	return nil
}

func (a *jsiiProxy_Application) validateOnDeploymentStartParameters(eventDestination IEventDestination, options *ExtensionOptions) error {
	return nil
}

func (a *jsiiProxy_Application) validateOnDeploymentStepParameters(eventDestination IEventDestination, options *ExtensionOptions) error {
	return nil
}

func (a *jsiiProxy_Application) validatePreCreateHostedConfigurationVersionParameters(eventDestination IEventDestination, options *ExtensionOptions) error {
	return nil
}

func (a *jsiiProxy_Application) validatePreStartDeploymentParameters(eventDestination IEventDestination, options *ExtensionOptions) error {
	return nil
}

func validateApplication_AddAgentToEcsParameters(taskDef awsecs.TaskDefinition) error {
	return nil
}

func validateApplication_FromApplicationArnParameters(scope constructs.Construct, id *string, applicationArn *string) error {
	return nil
}

func validateApplication_FromApplicationIdParameters(scope constructs.Construct, id *string, applicationId *string) error {
	return nil
}

func validateApplication_GetLambdaLayerVersionArnParameters(region *string) error {
	return nil
}

func validateApplication_IsConstructParameters(x interface{}) error {
	return nil
}

func validateApplication_IsOwnedResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateApplication_IsResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func (j *jsiiProxy_Application) validateSetExtensibleParameters(val ExtensibleBase) error {
	return nil
}

func validateNewApplicationParameters(scope constructs.Construct, id *string, props *ApplicationProps) error {
	return nil
}

