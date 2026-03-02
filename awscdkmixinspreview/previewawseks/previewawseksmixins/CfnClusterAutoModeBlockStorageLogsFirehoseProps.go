package previewawseksmixins


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnClusterAutoModeBlockStorageLogsFirehoseProps := &CfnClusterAutoModeBlockStorageLogsFirehoseProps{
//   	OutputFormat: awscdkmixinspreview.Mixins.CfnClusterAutoModeBlockStorageLogsOutputFormat.Firehose_JSON,
//   	RecordFields: []CfnClusterAutoModeBlockStorageLogsRecordFields{
//   		awscdkmixinspreview.*Mixins.CfnClusterAutoModeBlockStorageLogsRecordFields_LEVEL,
//   	},
//   }
//
// Experimental.
type CfnClusterAutoModeBlockStorageLogsFirehoseProps struct {
	// Format for log output, options are json,plain,raw.
	// Experimental.
	OutputFormat CfnClusterAutoModeBlockStorageLogsOutputFormat_Firehose `field:"optional" json:"outputFormat" yaml:"outputFormat"`
	// Record fields that can be provided to a log delivery.
	// Experimental.
	RecordFields *[]CfnClusterAutoModeBlockStorageLogsRecordFields `field:"optional" json:"recordFields" yaml:"recordFields"`
}

