//go:build !no_runtime_type_checking

package awsec2

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func validateIpAddresses_AwsIpamAllocationParameters(props *AwsIpamProps) error {
	if props == nil {
		return fmt.Errorf("parameter props is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(props, func() string { return "parameter props" }); err != nil {
		return err
	}

	return nil
}

func validateIpAddresses_CidrParameters(cidrBlock *string) error {
	if cidrBlock == nil {
		return fmt.Errorf("parameter cidrBlock is required, but nil was provided")
	}

	return nil
}

