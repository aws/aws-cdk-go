package previewawsqbusinessmixins


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnApplicationSyncJobLogsFirehoseProps := &CfnApplicationSyncJobLogsFirehoseProps{
//   	OutputFormat: awscdkmixinspreview.Mixins.CfnApplicationSyncJobLogsOutputFormat.Firehose_PLAIN,
//   	RecordFields: []CfnApplicationSyncJobLogsRecordFields{
//   		awscdkmixinspreview.*Mixins.CfnApplicationSyncJobLogsRecordFields_TIMESTAMP,
//   	},
//   }
//
// Experimental.
type CfnApplicationSyncJobLogsFirehoseProps struct {
	// Format for log output, options are plain,json,raw.
	// Experimental.
	OutputFormat CfnApplicationSyncJobLogsOutputFormat_Firehose `field:"optional" json:"outputFormat" yaml:"outputFormat"`
	// Record fields that can be provided to a log delivery.
	// Experimental.
	RecordFields *[]CfnApplicationSyncJobLogsRecordFields `field:"optional" json:"recordFields" yaml:"recordFields"`
}

