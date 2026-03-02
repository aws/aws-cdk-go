package previewawsentityresolutionmixins


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnMatchingWorkflowWorkflowLogsDestProps := &CfnMatchingWorkflowWorkflowLogsDestProps{
//   	RecordFields: []CfnMatchingWorkflowWorkflowLogsRecordFields{
//   		awscdkmixinspreview.Mixins.CfnMatchingWorkflowWorkflowLogsRecordFields_RESOURCE_ARN,
//   	},
//   }
//
// Experimental.
type CfnMatchingWorkflowWorkflowLogsDestProps struct {
	// Record fields that can be provided to a log delivery.
	// Experimental.
	RecordFields *[]CfnMatchingWorkflowWorkflowLogsRecordFields `field:"optional" json:"recordFields" yaml:"recordFields"`
}

