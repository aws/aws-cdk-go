package previewawspcsmixins


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnClusterPcsSchedulerAuditLogsLogGroupProps := &CfnClusterPcsSchedulerAuditLogsLogGroupProps{
//   	OutputFormat: awscdkmixinspreview.Mixins.CfnClusterPcsSchedulerAuditLogsOutputFormat.LogGroup_PLAIN,
//   	RecordFields: []CfnClusterPcsSchedulerAuditLogsRecordFields{
//   		awscdkmixinspreview.*Mixins.CfnClusterPcsSchedulerAuditLogsRecordFields_RESOURCE_ID,
//   	},
//   }
//
// Experimental.
type CfnClusterPcsSchedulerAuditLogsLogGroupProps struct {
	// Format for log output, options are plain,json.
	// Experimental.
	OutputFormat CfnClusterPcsSchedulerAuditLogsOutputFormat_LogGroup `field:"optional" json:"outputFormat" yaml:"outputFormat"`
	// Record fields that can be provided to a log delivery.
	// Experimental.
	RecordFields *[]CfnClusterPcsSchedulerAuditLogsRecordFields `field:"optional" json:"recordFields" yaml:"recordFields"`
}

