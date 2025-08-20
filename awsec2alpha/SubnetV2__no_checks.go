//go:build no_runtime_type_checking

package awsec2alpha

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_SubnetV2) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (s *jsiiProxy_SubnetV2) validateAssociateNetworkAclParameters(id *string, networkAcl awsec2.INetworkAcl) error {
	return nil
}

func (s *jsiiProxy_SubnetV2) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (s *jsiiProxy_SubnetV2) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func validateSubnetV2_FromSubnetV2AttributesParameters(scope constructs.Construct, id *string, attrs *SubnetV2Attributes) error {
	return nil
}

func validateSubnetV2_IsConstructParameters(x interface{}) error {
	return nil
}

func validateSubnetV2_IsOwnedResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateSubnetV2_IsResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateNewSubnetV2Parameters(scope constructs.Construct, id *string, props *SubnetV2Props) error {
	return nil
}

