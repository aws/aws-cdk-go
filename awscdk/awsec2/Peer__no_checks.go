//go:build no_runtime_type_checking

package awsec2

// Building without runtime type checking enabled, so all the below just return nil

func validatePeer_Ipv4Parameters(cidrIp *string) error {
	return nil
}

func validatePeer_Ipv6Parameters(cidrIp *string) error {
	return nil
}

func validatePeer_PrefixListParameters(prefixListId *string) error {
	return nil
}

func validatePeer_SecurityGroupIdParameters(securityGroupId *string) error {
	return nil
}

