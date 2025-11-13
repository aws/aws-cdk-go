package interfacesawsimagebuilder


// A reference to a Workflow resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   workflowReference := &WorkflowReference{
//   	WorkflowArn: jsii.String("workflowArn"),
//   }
//
type WorkflowReference struct {
	// The Arn of the Workflow resource.
	WorkflowArn *string `field:"required" json:"workflowArn" yaml:"workflowArn"`
}

