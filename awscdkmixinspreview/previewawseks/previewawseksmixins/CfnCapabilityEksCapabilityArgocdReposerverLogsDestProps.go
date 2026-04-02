package previewawseksmixins


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnCapabilityEksCapabilityArgocdReposerverLogsDestProps := &CfnCapabilityEksCapabilityArgocdReposerverLogsDestProps{
//   	RecordFields: []CfnCapabilityEksCapabilityArgocdReposerverLogsRecordFields{
//   		awscdkmixinspreview.Mixins.CfnCapabilityEksCapabilityArgocdReposerverLogsRecordFields_STREAM,
//   	},
//   }
//
// Experimental.
type CfnCapabilityEksCapabilityArgocdReposerverLogsDestProps struct {
	// Record fields that can be provided to a log delivery.
	// Experimental.
	RecordFields *[]CfnCapabilityEksCapabilityArgocdReposerverLogsRecordFields `field:"optional" json:"recordFields" yaml:"recordFields"`
}

