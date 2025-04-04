//go:build no_runtime_type_checking

package awsec2

// Building without runtime type checking enabled, so all the below just return nil

func validateAclCidr_Ipv4Parameters(ipv4Cidr *string) error {
	return nil
}

func validateAclCidr_Ipv6Parameters(ipv6Cidr *string) error {
	return nil
}

