package previewawsecsmixins


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnClusterActionLogsFirehoseProps := &CfnClusterActionLogsFirehoseProps{
//   	OutputFormat: awscdkmixinspreview.Mixins.CfnClusterActionLogsOutputFormat.Firehose_JSON,
//   	RecordFields: []CfnClusterActionLogsRecordFields{
//   		awscdkmixinspreview.*Mixins.CfnClusterActionLogsRecordFields_TIMESTAMP,
//   	},
//   }
//
// Experimental.
type CfnClusterActionLogsFirehoseProps struct {
	// Format for log output, options are json.
	// Experimental.
	OutputFormat CfnClusterActionLogsOutputFormat_Firehose `field:"optional" json:"outputFormat" yaml:"outputFormat"`
	// Record fields that can be provided to a log delivery.
	// Experimental.
	RecordFields *[]CfnClusterActionLogsRecordFields `field:"optional" json:"recordFields" yaml:"recordFields"`
}

