package awsconnect


// Contains information about a page configuration in a workspace, including the view assigned to the page.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   workspacePageProperty := &WorkspacePageProperty{
//   	Page: jsii.String("page"),
//   	ResourceArn: jsii.String("resourceArn"),
//
//   	// the properties below are optional
//   	InputData: jsii.String("inputData"),
//   	Slug: jsii.String("slug"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-workspace-workspacepage.html
//
type CfnWorkspace_WorkspacePageProperty struct {
	// The page identifier.
	//
	// System pages include `HOME` and `AGENT_EXPERIENCE` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-workspace-workspacepage.html#cfn-connect-workspace-workspacepage-page
	//
	Page *string `field:"required" json:"page" yaml:"page"`
	// The Amazon Resource Name (ARN) of the view associated with this page.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-workspace-workspacepage.html#cfn-connect-workspace-workspacepage-resourcearn
	//
	ResourceArn *string `field:"required" json:"resourceArn" yaml:"resourceArn"`
	// A JSON string containing input parameters passed to the view when the page is rendered.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-workspace-workspacepage.html#cfn-connect-workspace-workspacepage-inputdata
	//
	InputData *string `field:"optional" json:"inputData" yaml:"inputData"`
	// The URL-friendly identifier for the page.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-workspace-workspacepage.html#cfn-connect-workspace-workspacepage-slug
	//
	Slug *string `field:"optional" json:"slug" yaml:"slug"`
}

