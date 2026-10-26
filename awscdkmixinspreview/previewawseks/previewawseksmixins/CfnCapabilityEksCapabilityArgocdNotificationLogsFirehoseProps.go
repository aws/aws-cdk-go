package previewawseksmixins


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnCapabilityEksCapabilityArgocdNotificationLogsFirehoseProps := &CfnCapabilityEksCapabilityArgocdNotificationLogsFirehoseProps{
//   	OutputFormat: awscdkmixinspreview.Mixins.CfnCapabilityEksCapabilityArgocdNotificationLogsOutputFormat.Firehose_PLAIN,
//   	RecordFields: []CfnCapabilityEksCapabilityArgocdNotificationLogsRecordFields{
//   		awscdkmixinspreview.*Mixins.CfnCapabilityEksCapabilityArgocdNotificationLogsRecordFields_TIMESTAMP,
//   	},
//   }
//
// Experimental.
type CfnCapabilityEksCapabilityArgocdNotificationLogsFirehoseProps struct {
	// Format for log output, options are plain,json,raw.
	// Experimental.
	OutputFormat CfnCapabilityEksCapabilityArgocdNotificationLogsOutputFormat_Firehose `field:"optional" json:"outputFormat" yaml:"outputFormat"`
	// Record fields that can be provided to a log delivery.
	// Experimental.
	RecordFields *[]CfnCapabilityEksCapabilityArgocdNotificationLogsRecordFields `field:"optional" json:"recordFields" yaml:"recordFields"`
}

