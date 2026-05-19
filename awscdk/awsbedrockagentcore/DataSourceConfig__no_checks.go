//go:build no_runtime_type_checking

package awsbedrockagentcore

// Building without runtime type checking enabled, so all the below just return nil

func validateDataSourceConfig_FromAgentRuntimeEndpointParameters(runtime IBedrockAgentRuntime) error {
	return nil
}

func validateDataSourceConfig_FromAgentRuntimeEndpointNameParameters(runtime IBedrockAgentRuntime, endpointName *string) error {
	return nil
}

func validateDataSourceConfig_FromCloudWatchLogsParameters(config *CloudWatchLogsDataSourceConfig) error {
	return nil
}

