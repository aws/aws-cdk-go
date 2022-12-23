//go:build no_runtime_type_checking

package awsec2

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_SubnetNetworkAclAssociation) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (s *jsiiProxy_SubnetNetworkAclAssociation) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (s *jsiiProxy_SubnetNetworkAclAssociation) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func validateSubnetNetworkAclAssociation_FromSubnetNetworkAclAssociationAssociationIdParameters(scope constructs.Construct, id *string, subnetNetworkAclAssociationAssociationId *string) error {
	return nil
}

func validateSubnetNetworkAclAssociation_IsConstructParameters(x interface{}) error {
	return nil
}

func validateSubnetNetworkAclAssociation_IsOwnedResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateSubnetNetworkAclAssociation_IsResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateNewSubnetNetworkAclAssociationParameters(scope constructs.Construct, id *string, props *SubnetNetworkAclAssociationProps) error {
	return nil
}

