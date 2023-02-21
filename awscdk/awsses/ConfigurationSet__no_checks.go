//go:build no_runtime_type_checking

package awsses

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_ConfigurationSet) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (c *jsiiProxy_ConfigurationSet) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (c *jsiiProxy_ConfigurationSet) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func validateConfigurationSet_FromConfigurationSetNameParameters(scope constructs.Construct, id *string, configurationSetName *string) error {
	return nil
}

func validateConfigurationSet_IsConstructParameters(x interface{}) error {
	return nil
}

func validateConfigurationSet_IsOwnedResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateConfigurationSet_IsResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateNewConfigurationSetParameters(scope constructs.Construct, id *string, props *ConfigurationSetProps) error {
	return nil
}

