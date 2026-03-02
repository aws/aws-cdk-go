package previewawsmskmixins


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnReplicatorApplicationLogsFirehoseProps := &CfnReplicatorApplicationLogsFirehoseProps{
//   	OutputFormat: awscdkmixinspreview.Mixins.CfnReplicatorApplicationLogsOutputFormat.Firehose_JSON,
//   	RecordFields: []CfnReplicatorApplicationLogsRecordFields{
//   		awscdkmixinspreview.*Mixins.CfnReplicatorApplicationLogsRecordFields_REPLICATOR_NAME,
//   	},
//   }
//
// Experimental.
type CfnReplicatorApplicationLogsFirehoseProps struct {
	// Format for log output, options are json,plain,raw.
	// Experimental.
	OutputFormat CfnReplicatorApplicationLogsOutputFormat_Firehose `field:"optional" json:"outputFormat" yaml:"outputFormat"`
	// Record fields that can be provided to a log delivery.
	// Experimental.
	RecordFields *[]CfnReplicatorApplicationLogsRecordFields `field:"optional" json:"recordFields" yaml:"recordFields"`
}

