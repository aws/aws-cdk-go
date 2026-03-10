package previewawsbedrockagentcoremixins


// The configuration that specifies where to read agent traces for online evaluation.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   dataSourceConfigProperty := &DataSourceConfigProperty{
//   	CloudWatchLogs: &CloudWatchLogsInputConfigProperty{
//   		LogGroupNames: []*string{
//   			jsii.String("logGroupNames"),
//   		},
//   		ServiceNames: []*string{
//   			jsii.String("serviceNames"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-onlineevaluationconfig-datasourceconfig.html
//
type CfnOnlineEvaluationConfigPropsMixin_DataSourceConfigProperty struct {
	// The configuration for reading agent traces from CloudWatch logs.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-onlineevaluationconfig-datasourceconfig.html#cfn-bedrockagentcore-onlineevaluationconfig-datasourceconfig-cloudwatchlogs
	//
	CloudWatchLogs interface{} `field:"optional" json:"cloudWatchLogs" yaml:"cloudWatchLogs"`
}

