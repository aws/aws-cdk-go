package previewawseksmixins


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnCapabilityEksCapabilityArgocdCommitserverLogsFirehoseProps := &CfnCapabilityEksCapabilityArgocdCommitserverLogsFirehoseProps{
//   	OutputFormat: awscdkmixinspreview.Mixins.CfnCapabilityEksCapabilityArgocdCommitserverLogsOutputFormat.Firehose_PLAIN,
//   	RecordFields: []CfnCapabilityEksCapabilityArgocdCommitserverLogsRecordFields{
//   		awscdkmixinspreview.*Mixins.CfnCapabilityEksCapabilityArgocdCommitserverLogsRecordFields_TIMESTAMP,
//   	},
//   }
//
// Experimental.
type CfnCapabilityEksCapabilityArgocdCommitserverLogsFirehoseProps struct {
	// Format for log output, options are plain,json,raw.
	// Experimental.
	OutputFormat CfnCapabilityEksCapabilityArgocdCommitserverLogsOutputFormat_Firehose `field:"optional" json:"outputFormat" yaml:"outputFormat"`
	// Record fields that can be provided to a log delivery.
	// Experimental.
	RecordFields *[]CfnCapabilityEksCapabilityArgocdCommitserverLogsRecordFields `field:"optional" json:"recordFields" yaml:"recordFields"`
}

