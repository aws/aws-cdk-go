package awstransfer


// Specifies the workflow ID for the workflow to assign and the execution role used for executing the workflow.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   workflowDetailProperty := &workflowDetailProperty{
//   	executionRole: jsii.String("executionRole"),
//   	workflowId: jsii.String("workflowId"),
//   }
//
type CfnServer_WorkflowDetailProperty struct {
	// Includes the necessary permissions for S3, EFS, and Lambda operations that Transfer can assume, so that all workflow steps can operate on the required resources.
	ExecutionRole *string `field:"required" json:"executionRole" yaml:"executionRole"`
	// A unique identifier for the workflow.
	WorkflowId *string `field:"required" json:"workflowId" yaml:"workflowId"`
}

