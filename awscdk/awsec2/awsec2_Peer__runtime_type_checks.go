//go:build !no_runtime_type_checking

package awsec2

import (
	"fmt"
)

func validatePeer_Ipv4Parameters(cidrIp *string) error {
	if cidrIp == nil {
		return fmt.Errorf("parameter cidrIp is required, but nil was provided")
	}

	return nil
}

func validatePeer_Ipv6Parameters(cidrIp *string) error {
	if cidrIp == nil {
		return fmt.Errorf("parameter cidrIp is required, but nil was provided")
	}

	return nil
}

func validatePeer_PrefixListParameters(prefixListId *string) error {
	if prefixListId == nil {
		return fmt.Errorf("parameter prefixListId is required, but nil was provided")
	}

	return nil
}

func validatePeer_SecurityGroupIdParameters(securityGroupId *string) error {
	if securityGroupId == nil {
		return fmt.Errorf("parameter securityGroupId is required, but nil was provided")
	}

	return nil
}

