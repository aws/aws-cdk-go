//go:build no_runtime_type_checking

package awsec2alpha

// Building without runtime type checking enabled, so all the below just return nil

func (i *jsiiProxy_ITransitGatewayVpcAttachment) validateAddSubnetsParameters(subnets *[]awsec2.ISubnet) error {
	return nil
}

func (i *jsiiProxy_ITransitGatewayVpcAttachment) validateRemoveSubnetsParameters(subnets *[]awsec2.ISubnet) error {
	return nil
}

