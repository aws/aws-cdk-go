package awsimagebuilderalpha


// Properties for an EC2 Image Builder Workflow.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import imagebuilder_alpha "github.com/aws/aws-cdk-go/awsimagebuilderalpha"
//
//   workflowAttributes := &WorkflowAttributes{
//   	WorkflowArn: jsii.String("workflowArn"),
//   	WorkflowName: jsii.String("workflowName"),
//   	WorkflowType: imagebuilder_alpha.WorkflowType_BUILD,
//   	WorkflowVersion: jsii.String("workflowVersion"),
//   }
//
// Experimental.
type WorkflowAttributes struct {
	// The ARN of the workflow.
	// Default: - the ARN is automatically constructed if a workflowName and workflowType is provided, otherwise a
	// workflowArn is required.
	//
	// Experimental.
	WorkflowArn *string `field:"optional" json:"workflowArn" yaml:"workflowArn"`
	// The name of the workflow.
	// Default: - the name is automatically constructed if a workflowArn is provided, otherwise a workflowName is required.
	//
	// Experimental.
	WorkflowName *string `field:"optional" json:"workflowName" yaml:"workflowName"`
	// The type of the workflow.
	// Default: - the type is automatically constructed if a workflowArn is provided, otherwise a workflowType is required.
	//
	// Experimental.
	WorkflowType WorkflowType `field:"optional" json:"workflowType" yaml:"workflowType"`
	// The version of the workflow.
	// Default: x.x.x
	//
	// Experimental.
	WorkflowVersion *string `field:"optional" json:"workflowVersion" yaml:"workflowVersion"`
}

