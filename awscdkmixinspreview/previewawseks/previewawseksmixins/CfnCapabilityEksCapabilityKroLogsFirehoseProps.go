package previewawseksmixins


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnCapabilityEksCapabilityKroLogsFirehoseProps := &CfnCapabilityEksCapabilityKroLogsFirehoseProps{
//   	OutputFormat: awscdkmixinspreview.Mixins.CfnCapabilityEksCapabilityKroLogsOutputFormat.Firehose_JSON,
//   	RecordFields: []CfnCapabilityEksCapabilityKroLogsRecordFields{
//   		awscdkmixinspreview.*Mixins.CfnCapabilityEksCapabilityKroLogsRecordFields_STREAM,
//   	},
//   }
//
// Experimental.
type CfnCapabilityEksCapabilityKroLogsFirehoseProps struct {
	// Format for log output, options are json,plain,raw.
	// Experimental.
	OutputFormat CfnCapabilityEksCapabilityKroLogsOutputFormat_Firehose `field:"optional" json:"outputFormat" yaml:"outputFormat"`
	// Record fields that can be provided to a log delivery.
	// Experimental.
	RecordFields *[]CfnCapabilityEksCapabilityKroLogsRecordFields `field:"optional" json:"recordFields" yaml:"recordFields"`
}

