//go:build no_runtime_type_checking

package awslambda

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CapacityProvider) validateAddFunctionParameters(func_ IFunction, options *CapacityProviderFunctionOptions) error {
	return nil
}

func (c *jsiiProxy_CapacityProvider) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (c *jsiiProxy_CapacityProvider) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (c *jsiiProxy_CapacityProvider) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func validateCapacityProvider_FromCapacityProviderArnParameters(scope constructs.Construct, id *string, capacityProviderArn *string) error {
	return nil
}

func validateCapacityProvider_FromCapacityProviderAttributesParameters(scope constructs.Construct, id *string, attrs *CapacityProviderAttributes) error {
	return nil
}

func validateCapacityProvider_FromCapacityProviderNameParameters(scope constructs.Construct, id *string, capacityProviderName *string) error {
	return nil
}

func validateCapacityProvider_IsConstructParameters(x interface{}) error {
	return nil
}

func validateCapacityProvider_IsOwnedResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateCapacityProvider_IsResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateNewCapacityProviderParameters(scope constructs.Construct, id *string, props *CapacityProviderProps) error {
	return nil
}

