//go:build !no_runtime_type_checking

package awsec2

import (
	"fmt"
)

func validateAclCidr_Ipv4Parameters(ipv4Cidr *string) error {
	if ipv4Cidr == nil {
		return fmt.Errorf("parameter ipv4Cidr is required, but nil was provided")
	}

	return nil
}

func validateAclCidr_Ipv6Parameters(ipv6Cidr *string) error {
	if ipv6Cidr == nil {
		return fmt.Errorf("parameter ipv6Cidr is required, but nil was provided")
	}

	return nil
}

