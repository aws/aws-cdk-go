package previewawsdsqlmixins


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnClusterAuditLogsDestProps := &CfnClusterAuditLogsDestProps{
//   	RecordFields: []CfnClusterAuditLogsRecordFields{
//   		awscdkmixinspreview.Mixins.CfnClusterAuditLogsRecordFields_TIMESTAMP,
//   	},
//   }
//
// Experimental.
type CfnClusterAuditLogsDestProps struct {
	// Record fields that can be provided to a log delivery.
	// Experimental.
	RecordFields *[]CfnClusterAuditLogsRecordFields `field:"optional" json:"recordFields" yaml:"recordFields"`
}

