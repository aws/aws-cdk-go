//go:build no_runtime_type_checking
// +build no_runtime_type_checking

package awsec2

// Building without runtime type checking enabled, so all the below just return nil

func (p *jsiiProxy_PublicSubnet) validateAddDefaultInternetRouteParameters(gatewayId *string, gatewayAttachment awscdk.IDependable) error {
	return nil
}

func (p *jsiiProxy_PublicSubnet) validateAddDefaultNatRouteParameters(natGatewayId *string) error {
	return nil
}

func (p *jsiiProxy_PublicSubnet) validateAddRouteParameters(id *string, options *AddRouteOptions) error {
	return nil
}

func (p *jsiiProxy_PublicSubnet) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (p *jsiiProxy_PublicSubnet) validateAssociateNetworkAclParameters(id *string, networkAcl INetworkAcl) error {
	return nil
}

func (p *jsiiProxy_PublicSubnet) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (p *jsiiProxy_PublicSubnet) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func (p *jsiiProxy_PublicSubnet) validateOnSynthesizeParameters(session constructs.ISynthesisSession) error {
	return nil
}

func (p *jsiiProxy_PublicSubnet) validateSynthesizeParameters(session awscdk.ISynthesisSession) error {
	return nil
}

func validatePublicSubnet_FromPublicSubnetAttributesParameters(scope constructs.Construct, id *string, attrs *PublicSubnetAttributes) error {
	return nil
}

func validatePublicSubnet_FromSubnetAttributesParameters(scope constructs.Construct, id *string, attrs *SubnetAttributes) error {
	return nil
}

func validatePublicSubnet_FromSubnetIdParameters(scope constructs.Construct, id *string, subnetId *string) error {
	return nil
}

func validatePublicSubnet_IsConstructParameters(x interface{}) error {
	return nil
}

func validatePublicSubnet_IsResourceParameters(construct awscdk.IConstruct) error {
	return nil
}

func validatePublicSubnet_IsVpcSubnetParameters(x interface{}) error {
	return nil
}

func validateNewPublicSubnetParameters(scope constructs.Construct, id *string, props *PublicSubnetProps) error {
	return nil
}

