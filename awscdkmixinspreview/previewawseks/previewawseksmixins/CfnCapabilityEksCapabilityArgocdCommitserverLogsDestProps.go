package previewawseksmixins


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnCapabilityEksCapabilityArgocdCommitserverLogsDestProps := &CfnCapabilityEksCapabilityArgocdCommitserverLogsDestProps{
//   	RecordFields: []CfnCapabilityEksCapabilityArgocdCommitserverLogsRecordFields{
//   		awscdkmixinspreview.Mixins.CfnCapabilityEksCapabilityArgocdCommitserverLogsRecordFields_STREAM,
//   	},
//   }
//
// Experimental.
type CfnCapabilityEksCapabilityArgocdCommitserverLogsDestProps struct {
	// Record fields that can be provided to a log delivery.
	// Experimental.
	RecordFields *[]CfnCapabilityEksCapabilityArgocdCommitserverLogsRecordFields `field:"optional" json:"recordFields" yaml:"recordFields"`
}

