//go:build !no_runtime_type_checking

package awselasticloadbalancingv2

import (
	"fmt"
)

func validateSourceNatIpv6Prefix_FromIpv6PrefixParameters(prefix *string) error {
	if prefix == nil {
		return fmt.Errorf("parameter prefix is required, but nil was provided")
	}

	return nil
}

func validateNewSourceNatIpv6PrefixParameters(prefix *string) error {
	if prefix == nil {
		return fmt.Errorf("parameter prefix is required, but nil was provided")
	}

	return nil
}

