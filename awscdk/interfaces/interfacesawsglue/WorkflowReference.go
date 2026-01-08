package interfacesawsglue


// A reference to a Workflow resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   workflowReference := &WorkflowReference{
//   	WorkflowName: jsii.String("workflowName"),
//   }
//
type WorkflowReference struct {
	// The Name of the Workflow resource.
	WorkflowName *string `field:"required" json:"workflowName" yaml:"workflowName"`
}

