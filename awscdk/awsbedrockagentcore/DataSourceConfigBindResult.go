package awsbedrockagentcore


// The result of binding a DataSourceConfig.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
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
type DataSourceConfigBindResult struct {
	// The CloudWatch Logs data source configuration.
	CloudWatchLogs *CloudWatchLogsDataSourceConfig `field:"required" json:"cloudWatchLogs" yaml:"cloudWatchLogs"`
}

