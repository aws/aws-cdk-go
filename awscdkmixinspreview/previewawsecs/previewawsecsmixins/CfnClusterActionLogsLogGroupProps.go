package previewawsecsmixins


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnClusterActionLogsLogGroupProps := &CfnClusterActionLogsLogGroupProps{
//   	OutputFormat: awscdkmixinspreview.Mixins.CfnClusterActionLogsOutputFormat.LogGroup_JSON,
//   	RecordFields: []CfnClusterActionLogsRecordFields{
//   		awscdkmixinspreview.*Mixins.CfnClusterActionLogsRecordFields_TIMESTAMP,
//   	},
//   }
//
// Experimental.
type CfnClusterActionLogsLogGroupProps struct {
	// Format for log output, options are json.
	// Experimental.
	OutputFormat CfnClusterActionLogsOutputFormat_LogGroup `field:"optional" json:"outputFormat" yaml:"outputFormat"`
	// Record fields that can be provided to a log delivery.
	// Experimental.
	RecordFields *[]CfnClusterActionLogsRecordFields `field:"optional" json:"recordFields" yaml:"recordFields"`
}

