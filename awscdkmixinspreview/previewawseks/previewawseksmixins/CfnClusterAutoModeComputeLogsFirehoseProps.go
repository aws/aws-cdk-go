package previewawseksmixins


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnClusterAutoModeComputeLogsFirehoseProps := &CfnClusterAutoModeComputeLogsFirehoseProps{
//   	OutputFormat: awscdkmixinspreview.Mixins.CfnClusterAutoModeComputeLogsOutputFormat.Firehose_JSON,
//   	RecordFields: []CfnClusterAutoModeComputeLogsRecordFields{
//   		awscdkmixinspreview.*Mixins.CfnClusterAutoModeComputeLogsRecordFields_LEVEL,
//   	},
//   }
//
// Experimental.
type CfnClusterAutoModeComputeLogsFirehoseProps struct {
	// Format for log output, options are json,plain,raw.
	// Experimental.
	OutputFormat CfnClusterAutoModeComputeLogsOutputFormat_Firehose `field:"optional" json:"outputFormat" yaml:"outputFormat"`
	// Record fields that can be provided to a log delivery.
	// Experimental.
	RecordFields *[]CfnClusterAutoModeComputeLogsRecordFields `field:"optional" json:"recordFields" yaml:"recordFields"`
}

