//go:build no_runtime_type_checking

package awsec2

// Building without runtime type checking enabled, so all the below just return nil

func validateClientVpnUserBasedAuthentication_ActiveDirectoryParameters(directoryId *string) error {
	return nil
}

func validateClientVpnUserBasedAuthentication_FederatedParameters(samlProvider interfacesawsiam.ISAMLProviderRef) error {
	return nil
}

