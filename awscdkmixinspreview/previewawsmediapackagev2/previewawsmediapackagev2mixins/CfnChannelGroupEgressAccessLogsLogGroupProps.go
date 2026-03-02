package previewawsmediapackagev2mixins


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnChannelGroupEgressAccessLogsLogGroupProps := &CfnChannelGroupEgressAccessLogsLogGroupProps{
//   	OutputFormat: awscdkmixinspreview.Mixins.CfnChannelGroupEgressAccessLogsOutputFormat.LogGroup_PLAIN,
//   	RecordFields: []CfnChannelGroupEgressAccessLogsRecordFields{
//   		awscdkmixinspreview.*Mixins.CfnChannelGroupEgressAccessLogsRecordFields_RESOURCE_ARN,
//   	},
//   }
//
// Experimental.
type CfnChannelGroupEgressAccessLogsLogGroupProps struct {
	// Format for log output, options are plain,json.
	// Experimental.
	OutputFormat CfnChannelGroupEgressAccessLogsOutputFormat_LogGroup `field:"optional" json:"outputFormat" yaml:"outputFormat"`
	// Record fields that can be provided to a log delivery.
	// Experimental.
	RecordFields *[]CfnChannelGroupEgressAccessLogsRecordFields `field:"optional" json:"recordFields" yaml:"recordFields"`
}

