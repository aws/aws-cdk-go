//go:build no_runtime_type_checking

package awsbatch

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_ConsumableResource) validateApplyCrossStackReferenceStrengthParameters(strength awscdk.ReferenceStrength) error {
	return nil
}

func (c *jsiiProxy_ConsumableResource) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (c *jsiiProxy_ConsumableResource) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (c *jsiiProxy_ConsumableResource) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func validateConsumableResource_FromConsumableResourceArnParameters(scope constructs.Construct, id *string, consumableResourceArn *string) error {
	return nil
}

func validateConsumableResource_IsConstructParameters(x interface{}) error {
	return nil
}

func validateConsumableResource_IsOwnedResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateConsumableResource_IsResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateNewConsumableResourceParameters(scope constructs.Construct, id *string, props *ConsumableResourceProps) error {
	return nil
}

