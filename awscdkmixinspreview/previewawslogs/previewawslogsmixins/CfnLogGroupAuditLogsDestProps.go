package previewawslogsmixins


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnLogGroupAuditLogsDestProps := &CfnLogGroupAuditLogsDestProps{
//   	RecordFields: []CfnLogGroupAuditLogsRecordFields{
//   		awscdkmixinspreview.Mixins.CfnLogGroupAuditLogsRecordFields_POLICYNAME,
//   	},
//   }
//
// Experimental.
type CfnLogGroupAuditLogsDestProps struct {
	// Record fields that can be provided to a log delivery.
	// Experimental.
	RecordFields *[]CfnLogGroupAuditLogsRecordFields `field:"optional" json:"recordFields" yaml:"recordFields"`
}

