//go:build no_runtime_type_checking

package awsec2

// Building without runtime type checking enabled, so all the below just return nil

func (n *jsiiProxy_NetworkAclEntry) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (n *jsiiProxy_NetworkAclEntry) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (n *jsiiProxy_NetworkAclEntry) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func validateNetworkAclEntry_IsConstructParameters(x interface{}) error {
	return nil
}

func validateNetworkAclEntry_IsOwnedResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateNetworkAclEntry_IsResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateNewNetworkAclEntryParameters(scope constructs.Construct, id *string, props *NetworkAclEntryProps) error {
	return nil
}

