package previewawseksmixins


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnCapabilityEksCapabilityArgocdApplicationsetLogsDestProps := &CfnCapabilityEksCapabilityArgocdApplicationsetLogsDestProps{
//   	RecordFields: []CfnCapabilityEksCapabilityArgocdApplicationsetLogsRecordFields{
//   		awscdkmixinspreview.Mixins.CfnCapabilityEksCapabilityArgocdApplicationsetLogsRecordFields_STREAM,
//   	},
//   }
//
// Experimental.
type CfnCapabilityEksCapabilityArgocdApplicationsetLogsDestProps struct {
	// Record fields that can be provided to a log delivery.
	// Experimental.
	RecordFields *[]CfnCapabilityEksCapabilityArgocdApplicationsetLogsRecordFields `field:"optional" json:"recordFields" yaml:"recordFields"`
}

