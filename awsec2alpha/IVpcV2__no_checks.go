//go:build no_runtime_type_checking

package awsec2alpha

// Building without runtime type checking enabled, so all the below just return nil

func (i *jsiiProxy_IVpcV2) validateAddEgressOnlyInternetGatewayParameters(options *EgressOnlyInternetGatewayOptions) error {
	return nil
}

func (i *jsiiProxy_IVpcV2) validateAddInternetGatewayParameters(options *InternetGatewayOptions) error {
	return nil
}

func (i *jsiiProxy_IVpcV2) validateAddNatGatewayParameters(options *NatGatewayOptions) error {
	return nil
}

func (i *jsiiProxy_IVpcV2) validateEnableVpnGatewayV2Parameters(options *VPNGatewayV2Options) error {
	return nil
}

