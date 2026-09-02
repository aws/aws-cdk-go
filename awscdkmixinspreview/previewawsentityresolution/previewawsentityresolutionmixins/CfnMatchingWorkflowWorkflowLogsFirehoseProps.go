package previewawsentityresolutionmixins


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnMatchingWorkflowWorkflowLogsFirehoseProps := &CfnMatchingWorkflowWorkflowLogsFirehoseProps{
//   	OutputFormat: awscdkmixinspreview.Mixins.CfnMatchingWorkflowWorkflowLogsOutputFormat.Firehose_PLAIN,
//   	RecordFields: []CfnMatchingWorkflowWorkflowLogsRecordFields{
//   		awscdkmixinspreview.*Mixins.CfnMatchingWorkflowWorkflowLogsRecordFields_TIMESTAMP,
//   	},
//   }
//
// Experimental.
type CfnMatchingWorkflowWorkflowLogsFirehoseProps struct {
	// Format for log output, options are plain,json,raw.
	// Experimental.
	OutputFormat CfnMatchingWorkflowWorkflowLogsOutputFormat_Firehose `field:"optional" json:"outputFormat" yaml:"outputFormat"`
	// Record fields that can be provided to a log delivery.
	// Experimental.
	RecordFields *[]CfnMatchingWorkflowWorkflowLogsRecordFields `field:"optional" json:"recordFields" yaml:"recordFields"`
}

