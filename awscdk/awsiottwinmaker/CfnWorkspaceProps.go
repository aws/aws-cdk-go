package awsiottwinmaker


// Properties for defining a `CfnWorkspace`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnWorkspaceProps := &CfnWorkspaceProps{
//   	Role: jsii.String("role"),
//   	S3Location: jsii.String("s3Location"),
//   	WorkspaceId: jsii.String("workspaceId"),
//
//   	// the properties below are optional
//   	Description: jsii.String("description"),
//   	Tags: map[string]*string{
//   		"tagsKey": jsii.String("tags"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iottwinmaker-workspace.html
//
type CfnWorkspaceProps struct {
	// The ARN of the execution role associated with the workspace.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iottwinmaker-workspace.html#cfn-iottwinmaker-workspace-role
	//
	Role interface{} `field:"required" json:"role" yaml:"role"`
	// The ARN of the S3 bucket where resources associated with the workspace are stored.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iottwinmaker-workspace.html#cfn-iottwinmaker-workspace-s3location
	//
	S3Location interface{} `field:"required" json:"s3Location" yaml:"s3Location"`
	// The ID of the workspace.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iottwinmaker-workspace.html#cfn-iottwinmaker-workspace-workspaceid
	//
	WorkspaceId *string `field:"required" json:"workspaceId" yaml:"workspaceId"`
	// The description of the workspace.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iottwinmaker-workspace.html#cfn-iottwinmaker-workspace-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Metadata that you can use to manage the workspace.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iottwinmaker-workspace.html#cfn-iottwinmaker-workspace-tags
	//
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
}

