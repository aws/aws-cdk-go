package awsentityresolution


// A reference to a MatchingWorkflow resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   matchingWorkflowReference := &MatchingWorkflowReference{
//   	WorkflowName: jsii.String("workflowName"),
//   }
//
type MatchingWorkflowReference struct {
	// The WorkflowName of the MatchingWorkflow resource.
	WorkflowName *string `field:"required" json:"workflowName" yaml:"workflowName"`
}

