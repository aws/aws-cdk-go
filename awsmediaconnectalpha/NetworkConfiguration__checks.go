//go:build !no_runtime_type_checking

package awsmediaconnectalpha

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func validateNetworkConfiguration_PublicNetworkParameters(allowlistCidr *string) error {
	if allowlistCidr == nil {
		return fmt.Errorf("parameter allowlistCidr is required, but nil was provided")
	}

	return nil
}

func validateNetworkConfiguration_VpcParameters(vpcInterface *VpcInterfaceConfig) error {
	if vpcInterface == nil {
		return fmt.Errorf("parameter vpcInterface is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(vpcInterface, func() string { return "parameter vpcInterface" }); err != nil {
		return err
	}

	return nil
}

