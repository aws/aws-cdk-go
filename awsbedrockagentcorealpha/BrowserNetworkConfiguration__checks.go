//go:build !no_runtime_type_checking

package awsbedrockagentcorealpha

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/constructs-go/constructs/v10"
)

func validateBrowserNetworkConfiguration_UsingVpcParameters(scope constructs.Construct, vpcConfig *VpcConfigProps) error {
	if scope == nil {
		return fmt.Errorf("parameter scope is required, but nil was provided")
	}

	if vpcConfig == nil {
		return fmt.Errorf("parameter vpcConfig is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(vpcConfig, func() string { return "parameter vpcConfig" }); err != nil {
		return err
	}

	return nil
}

func validateNewBrowserNetworkConfigurationParameters(mode *string, vpcConfig *VpcConfigProps) error {
	if mode == nil {
		return fmt.Errorf("parameter mode is required, but nil was provided")
	}

	if err := _jsii_.ValidateStruct(vpcConfig, func() string { return "parameter vpcConfig" }); err != nil {
		return err
	}

	return nil
}

