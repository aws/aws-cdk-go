package awsbedrockagentcorealpha


// The result of binding a DataSourceConfig.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import bedrock_agentcore_alpha "github.com/aws/aws-cdk-go/awsbedrockagentcorealpha"
//
//   dataSourceConfigBindResult := &DataSourceConfigBindResult{
//   	CloudWatchLogs: &CloudWatchLogsDataSourceConfig{
//   		LogGroupNames: []*string{
//   			jsii.String("logGroupNames"),
//   		},
//   		ServiceNames: []*string{
//   			jsii.String("serviceNames"),
//   		},
//   	},
//   }
//
// Deprecated: Use the equivalent construct from `aws-cdk-lib/aws-bedrockagentcore` instead.
type DataSourceConfigBindResult struct {
	// The CloudWatch Logs data source configuration.
	// Deprecated: Use the equivalent construct from `aws-cdk-lib/aws-bedrockagentcore` instead.
	CloudWatchLogs *CloudWatchLogsDataSourceConfig `field:"required" json:"cloudWatchLogs" yaml:"cloudWatchLogs"`
}

