package previewawsentityresolutionmixins


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnIdMappingWorkflowWorkflowLogsFirehoseProps := &CfnIdMappingWorkflowWorkflowLogsFirehoseProps{
//   	OutputFormat: awscdkmixinspreview.Mixins.CfnIdMappingWorkflowWorkflowLogsOutputFormat.Firehose_PLAIN,
//   	RecordFields: []CfnIdMappingWorkflowWorkflowLogsRecordFields{
//   		awscdkmixinspreview.*Mixins.CfnIdMappingWorkflowWorkflowLogsRecordFields_TIMESTAMP,
//   	},
//   }
//
// Experimental.
type CfnIdMappingWorkflowWorkflowLogsFirehoseProps struct {
	// Format for log output, options are plain,json,raw.
	// Experimental.
	OutputFormat CfnIdMappingWorkflowWorkflowLogsOutputFormat_Firehose `field:"optional" json:"outputFormat" yaml:"outputFormat"`
	// Record fields that can be provided to a log delivery.
	// Experimental.
	RecordFields *[]CfnIdMappingWorkflowWorkflowLogsRecordFields `field:"optional" json:"recordFields" yaml:"recordFields"`
}

