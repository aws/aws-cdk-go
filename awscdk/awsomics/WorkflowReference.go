package awsomics


// A reference to a Workflow resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   workflowReference := &WorkflowReference{
//   	WorkflowArn: jsii.String("workflowArn"),
//   	WorkflowId: jsii.String("workflowId"),
//   }
//
type WorkflowReference struct {
	// The ARN of the Workflow resource.
	WorkflowArn *string `field:"required" json:"workflowArn" yaml:"workflowArn"`
	// The Id of the Workflow resource.
	WorkflowId *string `field:"required" json:"workflowId" yaml:"workflowId"`
}

