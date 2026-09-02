package previewawsmediapackagev2mixins


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnChannelGroupIngressAccessLogsFirehoseProps := &CfnChannelGroupIngressAccessLogsFirehoseProps{
//   	OutputFormat: awscdkmixinspreview.Mixins.CfnChannelGroupIngressAccessLogsOutputFormat.Firehose_PLAIN,
//   	RecordFields: []CfnChannelGroupIngressAccessLogsRecordFields{
//   		awscdkmixinspreview.*Mixins.CfnChannelGroupIngressAccessLogsRecordFields_TIMESTAMP,
//   	},
//   }
//
// Experimental.
type CfnChannelGroupIngressAccessLogsFirehoseProps struct {
	// Format for log output, options are plain,json,raw.
	// Experimental.
	OutputFormat CfnChannelGroupIngressAccessLogsOutputFormat_Firehose `field:"optional" json:"outputFormat" yaml:"outputFormat"`
	// Record fields that can be provided to a log delivery.
	// Experimental.
	RecordFields *[]CfnChannelGroupIngressAccessLogsRecordFields `field:"optional" json:"recordFields" yaml:"recordFields"`
}

