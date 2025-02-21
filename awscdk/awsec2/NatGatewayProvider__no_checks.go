//go:build no_runtime_type_checking

package awsec2

// Building without runtime type checking enabled, so all the below just return nil

func (n *jsiiProxy_NatGatewayProvider) validateConfigureNatParameters(options *ConfigureNatOptions) error {
	return nil
}

func (n *jsiiProxy_NatGatewayProvider) validateConfigureSubnetParameters(subnet PrivateSubnet) error {
	return nil
}

func validateNatGatewayProvider_GatewayParameters(props *NatGatewayProps) error {
	return nil
}

func validateNatGatewayProvider_InstanceParameters(props *NatInstanceProps) error {
	return nil
}

func validateNatGatewayProvider_InstanceV2Parameters(props *NatInstanceProps) error {
	return nil
}

func validateNewNatGatewayProviderParameters(props *NatGatewayProps) error {
	return nil
}

