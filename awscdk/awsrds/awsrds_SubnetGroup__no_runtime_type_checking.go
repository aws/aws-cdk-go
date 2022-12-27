//go:build no_runtime_type_checking

package awsrds

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_SubnetGroup) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (s *jsiiProxy_SubnetGroup) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (s *jsiiProxy_SubnetGroup) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func validateSubnetGroup_FromSubnetGroupNameParameters(scope constructs.Construct, id *string, subnetGroupName *string) error {
	return nil
}

func validateSubnetGroup_IsConstructParameters(x interface{}) error {
	return nil
}

func validateSubnetGroup_IsOwnedResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateSubnetGroup_IsResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateNewSubnetGroupParameters(scope constructs.Construct, id *string, props *SubnetGroupProps) error {
	return nil
}

