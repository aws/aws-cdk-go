package previewawsentityresolutionmixins


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnMatchingWorkflowWorkflowLogsLogGroupProps := &CfnMatchingWorkflowWorkflowLogsLogGroupProps{
//   	OutputFormat: awscdkmixinspreview.Mixins.CfnMatchingWorkflowWorkflowLogsOutputFormat.LogGroup_PLAIN,
//   	RecordFields: []CfnMatchingWorkflowWorkflowLogsRecordFields{
//   		awscdkmixinspreview.*Mixins.CfnMatchingWorkflowWorkflowLogsRecordFields_RESOURCE_ARN,
//   	},
//   }
//
// Experimental.
type CfnMatchingWorkflowWorkflowLogsLogGroupProps struct {
	// Format for log output, options are plain,json.
	// Experimental.
	OutputFormat CfnMatchingWorkflowWorkflowLogsOutputFormat_LogGroup `field:"optional" json:"outputFormat" yaml:"outputFormat"`
	// Record fields that can be provided to a log delivery.
	// Experimental.
	RecordFields *[]CfnMatchingWorkflowWorkflowLogsRecordFields `field:"optional" json:"recordFields" yaml:"recordFields"`
}

