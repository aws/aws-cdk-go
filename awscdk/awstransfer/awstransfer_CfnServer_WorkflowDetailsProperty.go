package awstransfer


// Container for the `WorkflowDetail` data type.
//
// It is used by actions that trigger a workflow to begin execution.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   workflowDetailsProperty := &workflowDetailsProperty{
//   	onUpload: []interface{}{
//   		&workflowDetailProperty{
//   			executionRole: jsii.String("executionRole"),
//   			workflowId: jsii.String("workflowId"),
//   		},
//   	},
//   }
//
type CfnServer_WorkflowDetailsProperty struct {
	// A trigger that starts a workflow: the workflow begins to execute after a file is uploaded.
	//
	// To remove an associated workflow from a server, you can provide an empty `OnUpload` object, as in the following example.
	//
	// `aws transfer update-server --server-id s-01234567890abcdef --workflow-details '{"OnUpload":[]}'`.
	OnUpload interface{} `field:"required" json:"onUpload" yaml:"onUpload"`
}

