package awsentityresolution


// A reference to a IdMappingWorkflow resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   idMappingWorkflowReference := &IdMappingWorkflowReference{
//   	WorkflowName: jsii.String("workflowName"),
//   }
//
type IdMappingWorkflowReference struct {
	// The WorkflowName of the IdMappingWorkflow resource.
	WorkflowName *string `field:"required" json:"workflowName" yaml:"workflowName"`
}

