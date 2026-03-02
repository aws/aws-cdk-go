package previewawspcsmixins


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnClusterPcsSchedulerLogsFirehoseProps := &CfnClusterPcsSchedulerLogsFirehoseProps{
//   	OutputFormat: awscdkmixinspreview.Mixins.CfnClusterPcsSchedulerLogsOutputFormat.Firehose_JSON,
//   	RecordFields: []CfnClusterPcsSchedulerLogsRecordFields{
//   		awscdkmixinspreview.*Mixins.CfnClusterPcsSchedulerLogsRecordFields_RESOURCE_ID,
//   	},
//   }
//
// Experimental.
type CfnClusterPcsSchedulerLogsFirehoseProps struct {
	// Format for log output, options are json,plain,raw.
	// Experimental.
	OutputFormat CfnClusterPcsSchedulerLogsOutputFormat_Firehose `field:"optional" json:"outputFormat" yaml:"outputFormat"`
	// Record fields that can be provided to a log delivery.
	// Experimental.
	RecordFields *[]CfnClusterPcsSchedulerLogsRecordFields `field:"optional" json:"recordFields" yaml:"recordFields"`
}

