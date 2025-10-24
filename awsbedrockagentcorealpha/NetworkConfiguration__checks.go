//go:build !no_runtime_type_checking

package awsbedrockagentcorealpha

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func validateNewNetworkConfigurationParameters(mode *string, vpcConfig *VpcConfigProps) error {
	if mode == nil {
		return fmt.Errorf("parameter mode is required, but nil was provided")
	}

	if err := _jsii_.ValidateStruct(vpcConfig, func() string { return "parameter vpcConfig" }); err != nil {
		return err
	}

	return nil
}

