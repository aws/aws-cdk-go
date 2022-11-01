//go:build !no_runtime_type_checking

package awsec2

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func validateAclTraffic_IcmpParameters(props *AclIcmp) error {
	if props == nil {
		return fmt.Errorf("parameter props is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(props, func() string { return "parameter props" }); err != nil {
		return err
	}

	return nil
}

func validateAclTraffic_Icmpv6Parameters(props *AclIcmp) error {
	if props == nil {
		return fmt.Errorf("parameter props is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(props, func() string { return "parameter props" }); err != nil {
		return err
	}

	return nil
}

func validateAclTraffic_TcpPortParameters(port *float64) error {
	if port == nil {
		return fmt.Errorf("parameter port is required, but nil was provided")
	}

	return nil
}

func validateAclTraffic_TcpPortRangeParameters(startPort *float64, endPort *float64) error {
	if startPort == nil {
		return fmt.Errorf("parameter startPort is required, but nil was provided")
	}

	if endPort == nil {
		return fmt.Errorf("parameter endPort is required, but nil was provided")
	}

	return nil
}

func validateAclTraffic_UdpPortParameters(port *float64) error {
	if port == nil {
		return fmt.Errorf("parameter port is required, but nil was provided")
	}

	return nil
}

func validateAclTraffic_UdpPortRangeParameters(startPort *float64, endPort *float64) error {
	if startPort == nil {
		return fmt.Errorf("parameter startPort is required, but nil was provided")
	}

	if endPort == nil {
		return fmt.Errorf("parameter endPort is required, but nil was provided")
	}

	return nil
}

