package previewawsentityresolutionmixins


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnIdMappingWorkflowWorkflowLogsDestProps := &CfnIdMappingWorkflowWorkflowLogsDestProps{
//   	RecordFields: []CfnIdMappingWorkflowWorkflowLogsRecordFields{
//   		awscdkmixinspreview.Mixins.CfnIdMappingWorkflowWorkflowLogsRecordFields_RESOURCE_ARN,
//   	},
//   }
//
// Experimental.
type CfnIdMappingWorkflowWorkflowLogsDestProps struct {
	// Record fields that can be provided to a log delivery.
	// Experimental.
	RecordFields *[]CfnIdMappingWorkflowWorkflowLogsRecordFields `field:"optional" json:"recordFields" yaml:"recordFields"`
}

