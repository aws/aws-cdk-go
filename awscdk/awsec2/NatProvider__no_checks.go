//go:build no_runtime_type_checking

package awsec2

// Building without runtime type checking enabled, so all the below just return nil

func (n *jsiiProxy_NatProvider) validateConfigureNatParameters(options *ConfigureNatOptions) error {
	return nil
}

func (n *jsiiProxy_NatProvider) validateConfigureSubnetParameters(subnet PrivateSubnet) error {
	return nil
}

func validateNatProvider_GatewayParameters(props *NatGatewayProps) error {
	return nil
}

func validateNatProvider_InstanceParameters(props *NatInstanceProps) error {
	return nil
}

