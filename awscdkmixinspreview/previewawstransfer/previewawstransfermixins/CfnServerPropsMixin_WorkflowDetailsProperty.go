package previewawstransfermixins


// Container for the `WorkflowDetail` data type.
//
// It is used by actions that trigger a workflow to begin execution.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   workflowDetailsProperty := &WorkflowDetailsProperty{
//   	OnPartialUpload: []interface{}{
//   		&WorkflowDetailProperty{
//   			ExecutionRole: jsii.String("executionRole"),
//   			WorkflowId: jsii.String("workflowId"),
//   		},
//   	},
//   	OnUpload: []interface{}{
//   		&WorkflowDetailProperty{
//   			ExecutionRole: jsii.String("executionRole"),
//   			WorkflowId: jsii.String("workflowId"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-transfer-server-workflowdetails.html
//
type CfnServerPropsMixin_WorkflowDetailsProperty struct {
	// A trigger that starts a workflow if a file is only partially uploaded.
	//
	// You can attach a workflow to a server that executes whenever there is a partial upload.
	//
	// A *partial upload* occurs when a file is open when the session disconnects.
	//
	// > `OnPartialUpload` can contain a maximum of one `WorkflowDetail` object.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-transfer-server-workflowdetails.html#cfn-transfer-server-workflowdetails-onpartialupload
	//
	OnPartialUpload interface{} `field:"optional" json:"onPartialUpload" yaml:"onPartialUpload"`
	// A trigger that starts a workflow: the workflow begins to execute after a file is uploaded.
	//
	// To remove an associated workflow from a server, you can provide an empty `OnUpload` object, as in the following example.
	//
	// `aws transfer update-server --server-id s-01234567890abcdef --workflow-details '{"OnUpload":[]}'`
	//
	// > `OnUpload` can contain a maximum of one `WorkflowDetail` object.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-transfer-server-workflowdetails.html#cfn-transfer-server-workflowdetails-onupload
	//
	OnUpload interface{} `field:"optional" json:"onUpload" yaml:"onUpload"`
}

