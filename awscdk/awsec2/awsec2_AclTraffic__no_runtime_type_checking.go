//go:build no_runtime_type_checking

package awsec2

// Building without runtime type checking enabled, so all the below just return nil

func validateAclTraffic_IcmpParameters(props *AclIcmp) error {
	return nil
}

func validateAclTraffic_Icmpv6Parameters(props *AclIcmp) error {
	return nil
}

func validateAclTraffic_TcpPortParameters(port *float64) error {
	return nil
}

func validateAclTraffic_TcpPortRangeParameters(startPort *float64, endPort *float64) error {
	return nil
}

func validateAclTraffic_UdpPortParameters(port *float64) error {
	return nil
}

func validateAclTraffic_UdpPortRangeParameters(startPort *float64, endPort *float64) error {
	return nil
}

