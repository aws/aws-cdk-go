//go:build no_runtime_type_checking

package awsses

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_ConfigurationSetEventDestination) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (c *jsiiProxy_ConfigurationSetEventDestination) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (c *jsiiProxy_ConfigurationSetEventDestination) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func validateConfigurationSetEventDestination_FromConfigurationSetEventDestinationIdParameters(scope constructs.Construct, id *string, configurationSetEventDestinationId *string) error {
	return nil
}

func validateConfigurationSetEventDestination_IsConstructParameters(x interface{}) error {
	return nil
}

func validateConfigurationSetEventDestination_IsOwnedResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateConfigurationSetEventDestination_IsResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateNewConfigurationSetEventDestinationParameters(scope constructs.Construct, id *string, props *ConfigurationSetEventDestinationProps) error {
	return nil
}

