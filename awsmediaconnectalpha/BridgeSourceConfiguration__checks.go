//go:build !no_runtime_type_checking

package awsmediaconnectalpha

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func validateBridgeSourceConfiguration_FlowParameters(source *BridgeFlowSource) error {
	if source == nil {
		return fmt.Errorf("parameter source is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(source, func() string { return "parameter source" }); err != nil {
		return err
	}

	return nil
}

func validateBridgeSourceConfiguration_NetworkParameters(source *BridgeNetworkSource) error {
	if source == nil {
		return fmt.Errorf("parameter source is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(source, func() string { return "parameter source" }); err != nil {
		return err
	}

	return nil
}

func validateNewBridgeSourceConfigurationParameters(flowConfig *BridgeFlowSource, networkConfig *BridgeNetworkSource) error {
	if err := _jsii_.ValidateStruct(flowConfig, func() string { return "parameter flowConfig" }); err != nil {
		return err
	}

	if err := _jsii_.ValidateStruct(networkConfig, func() string { return "parameter networkConfig" }); err != nil {
		return err
	}

	return nil
}

