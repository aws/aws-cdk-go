package previewawslogsmixins


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnLogGroupAuditLogsLogGroupProps := &CfnLogGroupAuditLogsLogGroupProps{
//   	OutputFormat: awscdkmixinspreview.Mixins.CfnLogGroupAuditLogsOutputFormat.LogGroup_PLAIN,
//   	RecordFields: []CfnLogGroupAuditLogsRecordFields{
//   		awscdkmixinspreview.*Mixins.CfnLogGroupAuditLogsRecordFields_POLICYNAME,
//   	},
//   }
//
// Experimental.
type CfnLogGroupAuditLogsLogGroupProps struct {
	// Format for log output, options are plain,json.
	// Experimental.
	OutputFormat CfnLogGroupAuditLogsOutputFormat_LogGroup `field:"optional" json:"outputFormat" yaml:"outputFormat"`
	// Record fields that can be provided to a log delivery.
	// Experimental.
	RecordFields *[]CfnLogGroupAuditLogsRecordFields `field:"optional" json:"recordFields" yaml:"recordFields"`
}

