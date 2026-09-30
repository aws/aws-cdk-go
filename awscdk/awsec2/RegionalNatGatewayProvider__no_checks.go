//go:build no_runtime_type_checking

package awsec2

// Building without runtime type checking enabled, so all the below just return nil

func (r *jsiiProxy_RegionalNatGatewayProvider) validateConfigureNatParameters(options *ConfigureNatOptions) error {
	return nil
}

func (r *jsiiProxy_RegionalNatGatewayProvider) validateConfigureSubnetParameters(subnet PrivateSubnet) error {
	return nil
}

func validateRegionalNatGatewayProvider_GatewayParameters(props *NatGatewayProps) error {
	return nil
}

func validateRegionalNatGatewayProvider_InstanceParameters(props *NatInstanceProps) error {
	return nil
}

func validateRegionalNatGatewayProvider_InstanceV2Parameters(props *NatInstanceProps) error {
	return nil
}

func validateRegionalNatGatewayProvider_RegionalGatewayParameters(props *RegionalNatGatewayProviderProps) error {
	return nil
}

func validateNewRegionalNatGatewayProviderParameters(props *RegionalNatGatewayProviderProps) error {
	return nil
}

