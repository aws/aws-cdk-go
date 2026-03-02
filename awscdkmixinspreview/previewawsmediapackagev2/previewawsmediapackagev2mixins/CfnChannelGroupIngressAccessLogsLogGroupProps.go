package previewawsmediapackagev2mixins


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnChannelGroupIngressAccessLogsLogGroupProps := &CfnChannelGroupIngressAccessLogsLogGroupProps{
//   	OutputFormat: awscdkmixinspreview.Mixins.CfnChannelGroupIngressAccessLogsOutputFormat.LogGroup_PLAIN,
//   	RecordFields: []CfnChannelGroupIngressAccessLogsRecordFields{
//   		awscdkmixinspreview.*Mixins.CfnChannelGroupIngressAccessLogsRecordFields_RESOURCE_ARN,
//   	},
//   }
//
// Experimental.
type CfnChannelGroupIngressAccessLogsLogGroupProps struct {
	// Format for log output, options are plain,json.
	// Experimental.
	OutputFormat CfnChannelGroupIngressAccessLogsOutputFormat_LogGroup `field:"optional" json:"outputFormat" yaml:"outputFormat"`
	// Record fields that can be provided to a log delivery.
	// Experimental.
	RecordFields *[]CfnChannelGroupIngressAccessLogsRecordFields `field:"optional" json:"recordFields" yaml:"recordFields"`
}

