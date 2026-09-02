package previewawseventsmixins


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnEventBusInfoLogsFirehoseProps := &CfnEventBusInfoLogsFirehoseProps{
//   	OutputFormat: awscdkmixinspreview.Mixins.CfnEventBusInfoLogsOutputFormat.Firehose_PLAIN,
//   	RecordFields: []CfnEventBusInfoLogsRecordFields{
//   		awscdkmixinspreview.*Mixins.CfnEventBusInfoLogsRecordFields_TIMESTAMP,
//   	},
//   }
//
// Experimental.
type CfnEventBusInfoLogsFirehoseProps struct {
	// Format for log output, options are plain,json,raw.
	// Experimental.
	OutputFormat CfnEventBusInfoLogsOutputFormat_Firehose `field:"optional" json:"outputFormat" yaml:"outputFormat"`
	// Record fields that can be provided to a log delivery.
	// Experimental.
	RecordFields *[]CfnEventBusInfoLogsRecordFields `field:"optional" json:"recordFields" yaml:"recordFields"`
}

