package previewawslogsmixins


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnLogGroupAuditLogsFirehoseProps := &CfnLogGroupAuditLogsFirehoseProps{
//   	OutputFormat: awscdkmixinspreview.Mixins.CfnLogGroupAuditLogsOutputFormat.Firehose_JSON,
//   	RecordFields: []CfnLogGroupAuditLogsRecordFields{
//   		awscdkmixinspreview.*Mixins.CfnLogGroupAuditLogsRecordFields_POLICYNAME,
//   	},
//   }
//
// Experimental.
type CfnLogGroupAuditLogsFirehoseProps struct {
	// Format for log output, options are json,plain,raw.
	// Experimental.
	OutputFormat CfnLogGroupAuditLogsOutputFormat_Firehose `field:"optional" json:"outputFormat" yaml:"outputFormat"`
	// Record fields that can be provided to a log delivery.
	// Experimental.
	RecordFields *[]CfnLogGroupAuditLogsRecordFields `field:"optional" json:"recordFields" yaml:"recordFields"`
}

