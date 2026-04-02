//go:build no_runtime_type_checking

package awsmediapackagev2alpha

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_ChannelPolicy) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (c *jsiiProxy_ChannelPolicy) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (c *jsiiProxy_ChannelPolicy) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func validateChannelPolicy_IsConstructParameters(x interface{}) error {
	return nil
}

func validateChannelPolicy_IsOwnedResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateChannelPolicy_IsResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateNewChannelPolicyParameters(scope constructs.Construct, id *string, props *ChannelPolicyProps) error {
	return nil
}

