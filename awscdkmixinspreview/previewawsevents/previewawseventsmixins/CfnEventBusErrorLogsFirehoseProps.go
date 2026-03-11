package previewawseventsmixins


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnEventBusErrorLogsFirehoseProps := &CfnEventBusErrorLogsFirehoseProps{
//   	OutputFormat: awscdkmixinspreview.Mixins.CfnEventBusErrorLogsOutputFormat.Firehose_JSON,
//   	RecordFields: []CfnEventBusErrorLogsRecordFields{
//   		awscdkmixinspreview.*Mixins.CfnEventBusErrorLogsRecordFields_RESOURCE_ARN,
//   	},
//   }
//
// Experimental.
type CfnEventBusErrorLogsFirehoseProps struct {
	// Format for log output, options are json,plain,raw.
	// Experimental.
	OutputFormat CfnEventBusErrorLogsOutputFormat_Firehose `field:"optional" json:"outputFormat" yaml:"outputFormat"`
	// Record fields that can be provided to a log delivery.
	// Experimental.
	RecordFields *[]CfnEventBusErrorLogsRecordFields `field:"optional" json:"recordFields" yaml:"recordFields"`
}

