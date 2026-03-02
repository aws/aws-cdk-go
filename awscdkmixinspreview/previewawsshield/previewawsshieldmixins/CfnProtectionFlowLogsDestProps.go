package previewawsshieldmixins


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnProtectionFlowLogsDestProps := &CfnProtectionFlowLogsDestProps{
//   	RecordFields: []CfnProtectionFlowLogsRecordFields{
//   		awscdkmixinspreview.Mixins.CfnProtectionFlowLogsRecordFields_SRCADDR,
//   	},
//   }
//
// Experimental.
type CfnProtectionFlowLogsDestProps struct {
	// Record fields that can be provided to a log delivery.
	// Experimental.
	RecordFields *[]CfnProtectionFlowLogsRecordFields `field:"optional" json:"recordFields" yaml:"recordFields"`
}

