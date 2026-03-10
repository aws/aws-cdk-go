package previewawspcsmixins


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnClusterPcsSchedulerAuditLogsFirehoseProps := &CfnClusterPcsSchedulerAuditLogsFirehoseProps{
//   	OutputFormat: awscdkmixinspreview.Mixins.CfnClusterPcsSchedulerAuditLogsOutputFormat.Firehose_JSON,
//   	RecordFields: []CfnClusterPcsSchedulerAuditLogsRecordFields{
//   		awscdkmixinspreview.*Mixins.CfnClusterPcsSchedulerAuditLogsRecordFields_RESOURCE_ID,
//   	},
//   }
//
// Experimental.
type CfnClusterPcsSchedulerAuditLogsFirehoseProps struct {
	// Format for log output, options are json,plain,raw.
	// Experimental.
	OutputFormat CfnClusterPcsSchedulerAuditLogsOutputFormat_Firehose `field:"optional" json:"outputFormat" yaml:"outputFormat"`
	// Record fields that can be provided to a log delivery.
	// Experimental.
	RecordFields *[]CfnClusterPcsSchedulerAuditLogsRecordFields `field:"optional" json:"recordFields" yaml:"recordFields"`
}

