//go:build !no_runtime_type_checking

package awsec2alpha

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func validateIpAddresses_AmazonProvidedIpv6Parameters(props *SecondaryAddressProps) error {
	if props == nil {
		return fmt.Errorf("parameter props is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(props, func() string { return "parameter props" }); err != nil {
		return err
	}

	return nil
}

func validateIpAddresses_Ipv4Parameters(ipv4Cidr *string, props *SecondaryAddressProps) error {
	if ipv4Cidr == nil {
		return fmt.Errorf("parameter ipv4Cidr is required, but nil was provided")
	}

	if err := _jsii_.ValidateStruct(props, func() string { return "parameter props" }); err != nil {
		return err
	}

	return nil
}

func validateIpAddresses_Ipv4IpamParameters(ipv4IpamOptions *IpamOptions) error {
	if ipv4IpamOptions == nil {
		return fmt.Errorf("parameter ipv4IpamOptions is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(ipv4IpamOptions, func() string { return "parameter ipv4IpamOptions" }); err != nil {
		return err
	}

	return nil
}

func validateIpAddresses_Ipv6IpamParameters(ipv6IpamOptions *IpamOptions) error {
	if ipv6IpamOptions == nil {
		return fmt.Errorf("parameter ipv6IpamOptions is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(ipv6IpamOptions, func() string { return "parameter ipv6IpamOptions" }); err != nil {
		return err
	}

	return nil
}

