package interfacesawsomics


// A reference to a WorkflowVersion resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   workflowVersionReference := &WorkflowVersionReference{
//   	WorkflowVersionArn: jsii.String("workflowVersionArn"),
//   }
//
type WorkflowVersionReference struct {
	// The Arn of the WorkflowVersion resource.
	WorkflowVersionArn *string `field:"required" json:"workflowVersionArn" yaml:"workflowVersionArn"`
}

