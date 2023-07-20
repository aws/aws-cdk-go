package awstransfer


// Specifies the workflow ID for the workflow to assign and the execution role that's used for executing the workflow.
//
// In addition to a workflow to execute when a file is uploaded completely, `WorkflowDetails` can also contain a workflow ID (and execution role) for a workflow to execute on partial upload. A partial upload occurs when a file is open when the session disconnects.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   workflowDetailProperty := &WorkflowDetailProperty{
//   	ExecutionRole: jsii.String("executionRole"),
//   	WorkflowId: jsii.String("workflowId"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-transfer-server-workflowdetail.html
//
type CfnServer_WorkflowDetailProperty struct {
	// Includes the necessary permissions for S3, EFS, and Lambda operations that Transfer can assume, so that all workflow steps can operate on the required resources.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-transfer-server-workflowdetail.html#cfn-transfer-server-workflowdetail-executionrole
	//
	ExecutionRole *string `field:"required" json:"executionRole" yaml:"executionRole"`
	// A unique identifier for the workflow.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-transfer-server-workflowdetail.html#cfn-transfer-server-workflowdetail-workflowid
	//
	WorkflowId *string `field:"required" json:"workflowId" yaml:"workflowId"`
}

