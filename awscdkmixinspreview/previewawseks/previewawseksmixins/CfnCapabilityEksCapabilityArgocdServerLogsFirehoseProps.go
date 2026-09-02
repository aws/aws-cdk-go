package previewawseksmixins


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnCapabilityEksCapabilityArgocdServerLogsFirehoseProps := &CfnCapabilityEksCapabilityArgocdServerLogsFirehoseProps{
//   	OutputFormat: awscdkmixinspreview.Mixins.CfnCapabilityEksCapabilityArgocdServerLogsOutputFormat.Firehose_PLAIN,
//   	RecordFields: []CfnCapabilityEksCapabilityArgocdServerLogsRecordFields{
//   		awscdkmixinspreview.*Mixins.CfnCapabilityEksCapabilityArgocdServerLogsRecordFields_TIMESTAMP,
//   	},
//   }
//
// Experimental.
type CfnCapabilityEksCapabilityArgocdServerLogsFirehoseProps struct {
	// Format for log output, options are plain,json,raw.
	// Experimental.
	OutputFormat CfnCapabilityEksCapabilityArgocdServerLogsOutputFormat_Firehose `field:"optional" json:"outputFormat" yaml:"outputFormat"`
	// Record fields that can be provided to a log delivery.
	// Experimental.
	RecordFields *[]CfnCapabilityEksCapabilityArgocdServerLogsRecordFields `field:"optional" json:"recordFields" yaml:"recordFields"`
}

