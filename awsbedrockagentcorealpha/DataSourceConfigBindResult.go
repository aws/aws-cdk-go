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
// Experimental.
type DataSourceConfigBindResult struct {
	// The CloudWatch Logs data source configuration.
	// Experimental.
	CloudWatchLogs *CloudWatchLogsDataSourceConfig `field:"required" json:"cloudWatchLogs" yaml:"cloudWatchLogs"`
}

