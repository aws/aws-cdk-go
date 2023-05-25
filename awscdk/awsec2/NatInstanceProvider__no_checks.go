//go:build no_runtime_type_checking

package awsec2

// Building without runtime type checking enabled, so all the below just return nil

func (n *jsiiProxy_NatInstanceProvider) validateConfigureNatParameters(options *ConfigureNatOptions) error {
	return nil
}

func (n *jsiiProxy_NatInstanceProvider) validateConfigureSubnetParameters(subnet PrivateSubnet) error {
	return nil
}

func validateNatInstanceProvider_GatewayParameters(props *NatGatewayProps) error {
	return nil
}

func validateNatInstanceProvider_InstanceParameters(props *NatInstanceProps) error {
	return nil
}

func validateNewNatInstanceProviderParameters(props *NatInstanceProps) error {
	return nil
}

