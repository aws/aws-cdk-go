package previewawspcsmixins


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnClusterPcsSchedulerAuditLogsDestProps := &CfnClusterPcsSchedulerAuditLogsDestProps{
//   	RecordFields: []CfnClusterPcsSchedulerAuditLogsRecordFields{
//   		awscdkmixinspreview.Mixins.CfnClusterPcsSchedulerAuditLogsRecordFields_RESOURCE_ID,
//   	},
//   }
//
// Experimental.
type CfnClusterPcsSchedulerAuditLogsDestProps struct {
	// Record fields that can be provided to a log delivery.
	// Experimental.
	RecordFields *[]CfnClusterPcsSchedulerAuditLogsRecordFields `field:"optional" json:"recordFields" yaml:"recordFields"`
}

