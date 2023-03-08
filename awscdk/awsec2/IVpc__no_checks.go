//go:build no_runtime_type_checking

package awsec2

// Building without runtime type checking enabled, so all the below just return nil

func (i *jsiiProxy_IVpc) validateAddClientVpnEndpointParameters(id *string, options *ClientVpnEndpointOptions) error {
	return nil
}

func (i *jsiiProxy_IVpc) validateAddFlowLogParameters(id *string, options *FlowLogOptions) error {
	return nil
}

func (i *jsiiProxy_IVpc) validateAddGatewayEndpointParameters(id *string, options *GatewayVpcEndpointOptions) error {
	return nil
}

func (i *jsiiProxy_IVpc) validateAddInterfaceEndpointParameters(id *string, options *InterfaceVpcEndpointOptions) error {
	return nil
}

func (i *jsiiProxy_IVpc) validateAddVpnConnectionParameters(id *string, options *VpnConnectionOptions) error {
	return nil
}

func (i *jsiiProxy_IVpc) validateEnableVpnGatewayParameters(options *EnableVpnGatewayOptions) error {
	return nil
}

func (i *jsiiProxy_IVpc) validateSelectSubnetsParameters(selection *SubnetSelection) error {
	return nil
}

