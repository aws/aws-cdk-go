package previewawseksmixins


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnCapabilityEksCapabilityArgocdReposerverLogsFirehoseProps := &CfnCapabilityEksCapabilityArgocdReposerverLogsFirehoseProps{
//   	OutputFormat: awscdkmixinspreview.Mixins.CfnCapabilityEksCapabilityArgocdReposerverLogsOutputFormat.Firehose_JSON,
//   	RecordFields: []CfnCapabilityEksCapabilityArgocdReposerverLogsRecordFields{
//   		awscdkmixinspreview.*Mixins.CfnCapabilityEksCapabilityArgocdReposerverLogsRecordFields_STREAM,
//   	},
//   }
//
// Experimental.
type CfnCapabilityEksCapabilityArgocdReposerverLogsFirehoseProps struct {
	// Format for log output, options are json,plain,raw.
	// Experimental.
	OutputFormat CfnCapabilityEksCapabilityArgocdReposerverLogsOutputFormat_Firehose `field:"optional" json:"outputFormat" yaml:"outputFormat"`
	// Record fields that can be provided to a log delivery.
	// Experimental.
	RecordFields *[]CfnCapabilityEksCapabilityArgocdReposerverLogsRecordFields `field:"optional" json:"recordFields" yaml:"recordFields"`
}

