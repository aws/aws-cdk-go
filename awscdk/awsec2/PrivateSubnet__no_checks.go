//go:build no_runtime_type_checking

package awsec2

// Building without runtime type checking enabled, so all the below just return nil

func (p *jsiiProxy_PrivateSubnet) validateAddDefaultInternetRouteParameters(gatewayId *string, gatewayAttachment constructs.IDependable) error {
	return nil
}

func (p *jsiiProxy_PrivateSubnet) validateAddDefaultNatRouteParameters(natGatewayId *string) error {
	return nil
}

func (p *jsiiProxy_PrivateSubnet) validateAddRouteParameters(id *string, options *AddRouteOptions) error {
	return nil
}

func (p *jsiiProxy_PrivateSubnet) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (p *jsiiProxy_PrivateSubnet) validateAssociateNetworkAclParameters(id *string, networkAcl INetworkAcl) error {
	return nil
}

func (p *jsiiProxy_PrivateSubnet) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (p *jsiiProxy_PrivateSubnet) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func validatePrivateSubnet_FromPrivateSubnetAttributesParameters(scope constructs.Construct, id *string, attrs *PrivateSubnetAttributes) error {
	return nil
}

func validatePrivateSubnet_FromSubnetAttributesParameters(scope constructs.Construct, id *string, attrs *SubnetAttributes) error {
	return nil
}

func validatePrivateSubnet_FromSubnetIdParameters(scope constructs.Construct, id *string, subnetId *string) error {
	return nil
}

func validatePrivateSubnet_IsConstructParameters(x interface{}) error {
	return nil
}

func validatePrivateSubnet_IsOwnedResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validatePrivateSubnet_IsResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validatePrivateSubnet_IsVpcSubnetParameters(x interface{}) error {
	return nil
}

func validateNewPrivateSubnetParameters(scope constructs.Construct, id *string, props *PrivateSubnetProps) error {
	return nil
}

