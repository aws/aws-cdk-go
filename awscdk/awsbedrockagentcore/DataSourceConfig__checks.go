//go:build !no_runtime_type_checking

package awsbedrockagentcore

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func validateDataSourceConfig_FromAgentRuntimeEndpointParameters(runtime IBedrockAgentRuntime) error {
	if runtime == nil {
		return fmt.Errorf("parameter runtime is required, but nil was provided")
	}

	return nil
}

func validateDataSourceConfig_FromAgentRuntimeEndpointNameParameters(runtime IBedrockAgentRuntime, endpointName *string) error {
	if runtime == nil {
		return fmt.Errorf("parameter runtime is required, but nil was provided")
	}

	if endpointName == nil {
		return fmt.Errorf("parameter endpointName is required, but nil was provided")
	}

	return nil
}

func validateDataSourceConfig_FromCloudWatchLogsParameters(config *CloudWatchLogsDataSourceConfig) error {
	if config == nil {
		return fmt.Errorf("parameter config is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(config, func() string { return "parameter config" }); err != nil {
		return err
	}

	return nil
}

