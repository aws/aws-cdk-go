//go:build no_runtime_type_checking

package awsec2alpha

// Building without runtime type checking enabled, so all the below just return nil

func validateIpAddresses_AmazonProvidedIpv6Parameters(props *SecondaryAddressProps) error {
	return nil
}

func validateIpAddresses_Ipv4Parameters(ipv4Cidr *string, props *SecondaryAddressProps) error {
	return nil
}

func validateIpAddresses_Ipv4IpamParameters(ipv4IpamOptions *IpamOptions) error {
	return nil
}

func validateIpAddresses_Ipv6ByoipPoolParameters(props *Ipv6PoolSecondaryAddressProps) error {
	return nil
}

func validateIpAddresses_Ipv6IpamParameters(ipv6IpamOptions *IpamOptions) error {
	return nil
}

