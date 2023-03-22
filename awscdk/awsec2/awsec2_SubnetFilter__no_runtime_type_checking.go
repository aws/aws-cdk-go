//go:build no_runtime_type_checking

package awsec2

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_SubnetFilter) validateSelectSubnetsParameters(_subnets *[]ISubnet) error {
	return nil
}

func validateSubnetFilter_AvailabilityZonesParameters(availabilityZones *[]*string) error {
	return nil
}

func validateSubnetFilter_ByCidrMaskParameters(mask *float64) error {
	return nil
}

func validateSubnetFilter_ByIdsParameters(subnetIds *[]*string) error {
	return nil
}

func validateSubnetFilter_ContainsIpAddressesParameters(ipv4addrs *[]*string) error {
	return nil
}

