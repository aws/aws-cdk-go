package previewawsmediapackagev2mixins


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnChannelGroupEgressAccessLogsFirehoseProps := &CfnChannelGroupEgressAccessLogsFirehoseProps{
//   	OutputFormat: awscdkmixinspreview.Mixins.CfnChannelGroupEgressAccessLogsOutputFormat.Firehose_JSON,
//   	RecordFields: []CfnChannelGroupEgressAccessLogsRecordFields{
//   		awscdkmixinspreview.*Mixins.CfnChannelGroupEgressAccessLogsRecordFields_RESOURCE_ARN,
//   	},
//   }
//
// Experimental.
type CfnChannelGroupEgressAccessLogsFirehoseProps struct {
	// Format for log output, options are json,plain,raw.
	// Experimental.
	OutputFormat CfnChannelGroupEgressAccessLogsOutputFormat_Firehose `field:"optional" json:"outputFormat" yaml:"outputFormat"`
	// Record fields that can be provided to a log delivery.
	// Experimental.
	RecordFields *[]CfnChannelGroupEgressAccessLogsRecordFields `field:"optional" json:"recordFields" yaml:"recordFields"`
}

