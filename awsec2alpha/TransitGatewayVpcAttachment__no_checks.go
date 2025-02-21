//go:build no_runtime_type_checking

package awsec2alpha

// Building without runtime type checking enabled, so all the below just return nil

func (t *jsiiProxy_TransitGatewayVpcAttachment) validateAddSubnetsParameters(subnets *[]awsec2.ISubnet) error {
	return nil
}

func (t *jsiiProxy_TransitGatewayVpcAttachment) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (t *jsiiProxy_TransitGatewayVpcAttachment) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (t *jsiiProxy_TransitGatewayVpcAttachment) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func (t *jsiiProxy_TransitGatewayVpcAttachment) validateRemoveSubnetsParameters(subnets *[]awsec2.ISubnet) error {
	return nil
}

func validateTransitGatewayVpcAttachment_IsConstructParameters(x interface{}) error {
	return nil
}

func validateTransitGatewayVpcAttachment_IsOwnedResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateTransitGatewayVpcAttachment_IsResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateNewTransitGatewayVpcAttachmentParameters(scope constructs.Construct, id *string, props *TransitGatewayVpcAttachmentProps) error {
	return nil
}

