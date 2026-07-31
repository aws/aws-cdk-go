//go:build no_runtime_type_checking

package awsmediaconnectalpha

// Building without runtime type checking enabled, so all the below just return nil

func validateNetworkConfiguration_PublicNetworkParameters(allowlistCidr *string) error {
	return nil
}

func validateNetworkConfiguration_VpcParameters(vpcInterface *VpcInterfaceConfig) error {
	return nil
}

