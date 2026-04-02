package previewawseksmixins


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnCapabilityEksCapabilityAckLogsDestProps := &CfnCapabilityEksCapabilityAckLogsDestProps{
//   	RecordFields: []CfnCapabilityEksCapabilityAckLogsRecordFields{
//   		awscdkmixinspreview.Mixins.CfnCapabilityEksCapabilityAckLogsRecordFields_STREAM,
//   	},
//   }
//
// Experimental.
type CfnCapabilityEksCapabilityAckLogsDestProps struct {
	// Record fields that can be provided to a log delivery.
	// Experimental.
	RecordFields *[]CfnCapabilityEksCapabilityAckLogsRecordFields `field:"optional" json:"recordFields" yaml:"recordFields"`
}

