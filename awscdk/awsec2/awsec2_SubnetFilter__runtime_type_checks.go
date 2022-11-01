//go:build !no_runtime_type_checking

package awsec2

import (
	"fmt"
)

func (s *jsiiProxy_SubnetFilter) validateSelectSubnetsParameters(_subnets *[]ISubnet) error {
	if _subnets == nil {
		return fmt.Errorf("parameter _subnets is required, but nil was provided")
	}

	return nil
}

func validateSubnetFilter_AvailabilityZonesParameters(availabilityZones *[]*string) error {
	if availabilityZones == nil {
		return fmt.Errorf("parameter availabilityZones is required, but nil was provided")
	}

	return nil
}

func validateSubnetFilter_ByCidrMaskParameters(mask *float64) error {
	if mask == nil {
		return fmt.Errorf("parameter mask is required, but nil was provided")
	}

	return nil
}

func validateSubnetFilter_ByIdsParameters(subnetIds *[]*string) error {
	if subnetIds == nil {
		return fmt.Errorf("parameter subnetIds is required, but nil was provided")
	}

	return nil
}

func validateSubnetFilter_ContainsIpAddressesParameters(ipv4addrs *[]*string) error {
	if ipv4addrs == nil {
		return fmt.Errorf("parameter ipv4addrs is required, but nil was provided")
	}

	return nil
}

