//go:build no_runtime_type_checking

package awsec2

// Building without runtime type checking enabled, so all the below just return nil

func (n *jsiiProxy_NatInstanceProviderV2) validateConfigureNatParameters(options *ConfigureNatOptions) error {
	return nil
}

func (n *jsiiProxy_NatInstanceProviderV2) validateConfigureSubnetParameters(subnet PrivateSubnet) error {
	return nil
}

func validateNatInstanceProviderV2_GatewayParameters(props *NatGatewayProps) error {
	return nil
}

func validateNatInstanceProviderV2_InstanceParameters(props *NatInstanceProps) error {
	return nil
}

func validateNatInstanceProviderV2_InstanceV2Parameters(props *NatInstanceProps) error {
	return nil
}

func validateNewNatInstanceProviderV2Parameters(props *NatInstanceProps) error {
	return nil
}

