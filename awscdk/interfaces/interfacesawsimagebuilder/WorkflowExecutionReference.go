package interfacesawsimagebuilder


// A reference to a WorkflowExecution resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   workflowExecutionReference := &WorkflowExecutionReference{
//   	WorkflowExecutionArn: jsii.String("workflowExecutionArn"),
//   }
//
type WorkflowExecutionReference struct {
	// The Arn of the WorkflowExecution resource.
	WorkflowExecutionArn *string `field:"required" json:"workflowExecutionArn" yaml:"workflowExecutionArn"`
}

