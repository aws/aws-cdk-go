package previewawsbedrockmixins


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnFlowApplicationLogsFirehoseProps := &CfnFlowApplicationLogsFirehoseProps{
//   	OutputFormat: awscdkmixinspreview.Mixins.CfnFlowApplicationLogsOutputFormat.Firehose_JSON,
//   	RecordFields: []CfnFlowApplicationLogsRecordFields{
//   		awscdkmixinspreview.*Mixins.CfnFlowApplicationLogsRecordFields_EVENT_TIMESTAMP,
//   	},
//   }
//
// Experimental.
type CfnFlowApplicationLogsFirehoseProps struct {
	// Format for log output, options are json,plain,raw.
	// Experimental.
	OutputFormat CfnFlowApplicationLogsOutputFormat_Firehose `field:"optional" json:"outputFormat" yaml:"outputFormat"`
	// Record fields that can be provided to a log delivery.
	// Experimental.
	RecordFields *[]CfnFlowApplicationLogsRecordFields `field:"optional" json:"recordFields" yaml:"recordFields"`
}

