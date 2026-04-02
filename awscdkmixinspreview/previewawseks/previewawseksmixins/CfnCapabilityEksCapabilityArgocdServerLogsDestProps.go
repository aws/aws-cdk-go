package previewawseksmixins


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnCapabilityEksCapabilityArgocdServerLogsDestProps := &CfnCapabilityEksCapabilityArgocdServerLogsDestProps{
//   	RecordFields: []CfnCapabilityEksCapabilityArgocdServerLogsRecordFields{
//   		awscdkmixinspreview.Mixins.CfnCapabilityEksCapabilityArgocdServerLogsRecordFields_STREAM,
//   	},
//   }
//
// Experimental.
type CfnCapabilityEksCapabilityArgocdServerLogsDestProps struct {
	// Record fields that can be provided to a log delivery.
	// Experimental.
	RecordFields *[]CfnCapabilityEksCapabilityArgocdServerLogsRecordFields `field:"optional" json:"recordFields" yaml:"recordFields"`
}

