package previewawsentityresolutionmixins


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnIdMappingWorkflowWorkflowLogsLogGroupProps := &CfnIdMappingWorkflowWorkflowLogsLogGroupProps{
//   	OutputFormat: awscdkmixinspreview.Mixins.CfnIdMappingWorkflowWorkflowLogsOutputFormat.LogGroup_PLAIN,
//   	RecordFields: []CfnIdMappingWorkflowWorkflowLogsRecordFields{
//   		awscdkmixinspreview.*Mixins.CfnIdMappingWorkflowWorkflowLogsRecordFields_RESOURCE_ARN,
//   	},
//   }
//
// Experimental.
type CfnIdMappingWorkflowWorkflowLogsLogGroupProps struct {
	// Format for log output, options are plain,json.
	// Experimental.
	OutputFormat CfnIdMappingWorkflowWorkflowLogsOutputFormat_LogGroup `field:"optional" json:"outputFormat" yaml:"outputFormat"`
	// Record fields that can be provided to a log delivery.
	// Experimental.
	RecordFields *[]CfnIdMappingWorkflowWorkflowLogsRecordFields `field:"optional" json:"recordFields" yaml:"recordFields"`
}

