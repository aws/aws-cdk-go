//go:build no_runtime_type_checking
// +build no_runtime_type_checking

package awsec2

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_Subnet) validateAddDefaultInternetRouteParameters(gatewayId *string, gatewayAttachment awscdk.IDependable) error {
	return nil
}

func (s *jsiiProxy_Subnet) validateAddDefaultNatRouteParameters(natGatewayId *string) error {
	return nil
}

func (s *jsiiProxy_Subnet) validateAddRouteParameters(id *string, options *AddRouteOptions) error {
	return nil
}

func (s *jsiiProxy_Subnet) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (s *jsiiProxy_Subnet) validateAssociateNetworkAclParameters(id *string, networkAcl INetworkAcl) error {
	return nil
}

func (s *jsiiProxy_Subnet) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (s *jsiiProxy_Subnet) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func (s *jsiiProxy_Subnet) validateOnSynthesizeParameters(session constructs.ISynthesisSession) error {
	return nil
}

func (s *jsiiProxy_Subnet) validateSynthesizeParameters(session awscdk.ISynthesisSession) error {
	return nil
}

func validateSubnet_FromSubnetAttributesParameters(scope constructs.Construct, id *string, attrs *SubnetAttributes) error {
	return nil
}

func validateSubnet_FromSubnetIdParameters(scope constructs.Construct, id *string, subnetId *string) error {
	return nil
}

func validateSubnet_IsConstructParameters(x interface{}) error {
	return nil
}

func validateSubnet_IsResourceParameters(construct awscdk.IConstruct) error {
	return nil
}

func validateSubnet_IsVpcSubnetParameters(x interface{}) error {
	return nil
}

func validateNewSubnetParameters(scope constructs.Construct, id *string, props *SubnetProps) error {
	return nil
}

