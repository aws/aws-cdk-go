//go:build no_runtime_type_checking

package awsmediaconnectalpha

// Building without runtime type checking enabled, so all the below just return nil

func validateRouterNetworkConfiguration_PublicNetworkParameters(props *PublicNetworkConfigurationProps) error {
	return nil
}

func validateRouterNetworkConfiguration_VpcParameters(props *VpcNetworkConfigurationProps) error {
	return nil
}

