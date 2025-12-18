package previewawsconnectmixins


// Contains information about a page configuration in a workspace, including the view assigned to the page.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   workspacePageProperty := &WorkspacePageProperty{
//   	InputData: jsii.String("inputData"),
//   	Page: jsii.String("page"),
//   	ResourceArn: jsii.String("resourceArn"),
//   	Slug: jsii.String("slug"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-workspace-workspacepage.html
//
type CfnWorkspacePropsMixin_WorkspacePageProperty struct {
	// A JSON string containing input parameters passed to the view when the page is rendered.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-workspace-workspacepage.html#cfn-connect-workspace-workspacepage-inputdata
	//
	InputData *string `field:"optional" json:"inputData" yaml:"inputData"`
	// The page identifier.
	//
	// System pages include `HOME` and `AGENT_EXPERIENCE` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-workspace-workspacepage.html#cfn-connect-workspace-workspacepage-page
	//
	Page *string `field:"optional" json:"page" yaml:"page"`
	// The Amazon Resource Name (ARN) of the view associated with this page.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-workspace-workspacepage.html#cfn-connect-workspace-workspacepage-resourcearn
	//
	ResourceArn *string `field:"optional" json:"resourceArn" yaml:"resourceArn"`
	// The URL-friendly identifier for the page.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-workspace-workspacepage.html#cfn-connect-workspace-workspacepage-slug
	//
	Slug *string `field:"optional" json:"slug" yaml:"slug"`
}

