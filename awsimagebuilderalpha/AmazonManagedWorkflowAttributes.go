package awsimagebuilderalpha


// Properties for an EC2 Image Builder Amazon-managed workflow.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import imagebuilder_alpha "github.com/aws/aws-cdk-go/awsimagebuilderalpha"
//
//   amazonManagedWorkflowAttributes := &AmazonManagedWorkflowAttributes{
//   	WorkflowName: jsii.String("workflowName"),
//   	WorkflowType: imagebuilder_alpha.WorkflowType_BUILD,
//   }
//
// Experimental.
type AmazonManagedWorkflowAttributes struct {
	// The name of the Amazon-managed workflow.
	// Experimental.
	WorkflowName *string `field:"required" json:"workflowName" yaml:"workflowName"`
	// The type of the Amazon-managed workflow.
	// Experimental.
	WorkflowType WorkflowType `field:"required" json:"workflowType" yaml:"workflowType"`
}

