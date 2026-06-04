package awsbedrockagentcorealpha

import (
	_init_ "github.com/aws/aws-cdk-go/awsbedrockagentcorealpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Configuration for the data source used in online evaluation.
//
// Use the static factory methods to create data source configurations:
// - `DataSourceConfig.fromAgentRuntimeEndpoint()` for AgentCore Runtime (recommended)
// - `DataSourceConfig.fromAgentRuntimeEndpointName()` for AgentCore Runtime using endpoint name string
// - `DataSourceConfig.fromCloudWatchLogs()` for external agents or custom log groups
//
// Example:
//   // CloudWatch Logs data source (for external agents)
//   dataSource := agentcore.DataSourceConfig_FromCloudWatchLogs(&CloudWatchLogsDataSourceConfig{
//   	LogGroupNames: []*string{
//   		jsii.String("/aws/my-external-agent/logs"),
//   	},
//   	ServiceNames: []*string{
//   		jsii.String("my-external-agent"),
//   	},
//   })
//
// Deprecated: Use the equivalent construct from `aws-cdk-lib/aws-bedrockagentcore` instead.
type DataSourceConfig interface {
	// The CloudWatch Logs configuration.
	// Deprecated: Use the equivalent construct from `aws-cdk-lib/aws-bedrockagentcore` instead.
	CloudWatchLogsConfig() *CloudWatchLogsDataSourceConfig
	// Binds the data source configuration to produce the L1 property.
	// Deprecated: Use the equivalent construct from `aws-cdk-lib/aws-bedrockagentcore` instead.
	Bind() *DataSourceConfigBindResult
}

// The jsii proxy struct for DataSourceConfig
type jsiiProxy_DataSourceConfig struct {
	_ byte // padding
}

func (j *jsiiProxy_DataSourceConfig) CloudWatchLogsConfig() *CloudWatchLogsDataSourceConfig {
	var returns *CloudWatchLogsDataSourceConfig
	_jsii_.Get(
		j,
		"cloudWatchLogsConfig",
		&returns,
	)
	return returns
}


// Creates a data source configuration from an AgentCore Runtime and optional endpoint.
//
// This is the recommended way to configure evaluation for AgentCore Runtime agents.
// It automatically derives the CloudWatch log group and service name from the runtime and endpoint.
//
// Returns: A DataSourceConfig instance.
//
// Example:
//   // Using a specific endpoint
//   var runtime Runtime
//
//   endpoint := runtime.AddEndpoint(jsii.String("PROD"))
//   dataSource := agentcore.DataSourceConfig_FromAgentRuntimeEndpoint(runtime, endpoint)
//
// Deprecated: Use the equivalent construct from `aws-cdk-lib/aws-bedrockagentcore` instead.
func DataSourceConfig_FromAgentRuntimeEndpoint(runtime IBedrockAgentRuntime, endpoint IRuntimeEndpoint) DataSourceConfig {
	_init_.Initialize()

	if err := validateDataSourceConfig_FromAgentRuntimeEndpointParameters(runtime); err != nil {
		panic(err)
	}
	var returns DataSourceConfig

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-bedrock-agentcore-alpha.DataSourceConfig",
		"fromAgentRuntimeEndpoint",
		[]interface{}{runtime, endpoint},
		&returns,
	)

	return returns
}

// Creates a data source configuration from an AgentCore Runtime and an endpoint name string.
//
// Use this method when you want to reference an endpoint by name without
// having a construct reference. For construct references, prefer `fromAgentRuntimeEndpoint()`.
//
// Returns: A DataSourceConfig instance.
//
// Example:
//   var runtime Runtime
//
//   dataSource := agentcore.DataSourceConfig_FromAgentRuntimeEndpointName(runtime, jsii.String("PROD"))
//
// Deprecated: Use the equivalent construct from `aws-cdk-lib/aws-bedrockagentcore` instead.
func DataSourceConfig_FromAgentRuntimeEndpointName(runtime IBedrockAgentRuntime, endpointName *string) DataSourceConfig {
	_init_.Initialize()

	if err := validateDataSourceConfig_FromAgentRuntimeEndpointNameParameters(runtime, endpointName); err != nil {
		panic(err)
	}
	var returns DataSourceConfig

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-bedrock-agentcore-alpha.DataSourceConfig",
		"fromAgentRuntimeEndpointName",
		[]interface{}{runtime, endpointName},
		&returns,
	)

	return returns
}

// Creates a CloudWatch Logs data source configuration.
//
// Use this when your agent traces are stored in CloudWatch Logs,
// such as for external agents or when you need to specify log groups directly.
//
// Returns: A DataSourceConfig instance.
//
// Example:
//   dataSource := agentcore.DataSourceConfig_FromCloudWatchLogs(&CloudWatchLogsDataSourceConfig{
//   	LogGroupNames: []*string{
//   		jsii.String("/aws/bedrock-agentcore/runtimes/myRuntime-abc123-DEFAULT"),
//   	},
//   	ServiceNames: []*string{
//   		jsii.String("myRuntime.DEFAULT"),
//   	},
//   })
//
// Deprecated: Use the equivalent construct from `aws-cdk-lib/aws-bedrockagentcore` instead.
func DataSourceConfig_FromCloudWatchLogs(config *CloudWatchLogsDataSourceConfig) DataSourceConfig {
	_init_.Initialize()

	if err := validateDataSourceConfig_FromCloudWatchLogsParameters(config); err != nil {
		panic(err)
	}
	var returns DataSourceConfig

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-bedrock-agentcore-alpha.DataSourceConfig",
		"fromCloudWatchLogs",
		[]interface{}{config},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataSourceConfig) Bind() *DataSourceConfigBindResult {
	var returns *DataSourceConfigBindResult

	_jsii_.Invoke(
		d,
		"bind",
		nil, // no parameters
		&returns,
	)

	return returns
}

