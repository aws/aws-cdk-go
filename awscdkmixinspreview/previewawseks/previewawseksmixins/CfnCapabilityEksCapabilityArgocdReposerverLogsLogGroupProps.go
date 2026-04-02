package previewawseksmixins


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnCapabilityEksCapabilityArgocdReposerverLogsLogGroupProps := &CfnCapabilityEksCapabilityArgocdReposerverLogsLogGroupProps{
//   	OutputFormat: awscdkmixinspreview.Mixins.CfnCapabilityEksCapabilityArgocdReposerverLogsOutputFormat.LogGroup_PLAIN,
//   	RecordFields: []CfnCapabilityEksCapabilityArgocdReposerverLogsRecordFields{
//   		awscdkmixinspreview.*Mixins.CfnCapabilityEksCapabilityArgocdReposerverLogsRecordFields_STREAM,
//   	},
//   }
//
// Experimental.
type CfnCapabilityEksCapabilityArgocdReposerverLogsLogGroupProps struct {
	// Format for log output, options are plain,json.
	// Experimental.
	OutputFormat CfnCapabilityEksCapabilityArgocdReposerverLogsOutputFormat_LogGroup `field:"optional" json:"outputFormat" yaml:"outputFormat"`
	// Record fields that can be provided to a log delivery.
	// Experimental.
	RecordFields *[]CfnCapabilityEksCapabilityArgocdReposerverLogsRecordFields `field:"optional" json:"recordFields" yaml:"recordFields"`
}

