package previewawseksmixins


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnCapabilityEksCapabilityKroLogsDestProps := &CfnCapabilityEksCapabilityKroLogsDestProps{
//   	RecordFields: []CfnCapabilityEksCapabilityKroLogsRecordFields{
//   		awscdkmixinspreview.Mixins.CfnCapabilityEksCapabilityKroLogsRecordFields_STREAM,
//   	},
//   }
//
// Experimental.
type CfnCapabilityEksCapabilityKroLogsDestProps struct {
	// Record fields that can be provided to a log delivery.
	// Experimental.
	RecordFields *[]CfnCapabilityEksCapabilityKroLogsRecordFields `field:"optional" json:"recordFields" yaml:"recordFields"`
}

