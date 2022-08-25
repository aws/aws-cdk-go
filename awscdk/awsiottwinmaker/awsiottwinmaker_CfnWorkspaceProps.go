package awsiottwinmaker


// Properties for defining a `CfnWorkspace`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnWorkspaceProps := &cfnWorkspaceProps{
//   	role: jsii.String("role"),
//   	s3Location: jsii.String("s3Location"),
//   	workspaceId: jsii.String("workspaceId"),
//
//   	// the properties below are optional
//   	description: jsii.String("description"),
//   	tags: map[string]*string{
//   		"tagsKey": jsii.String("tags"),
//   	},
//   }
//
type CfnWorkspaceProps struct {
	// The ARN of the execution role associated with the workspace.
	Role *string `field:"required" json:"role" yaml:"role"`
	// The ARN of the S3 bucket where resources associated with the workspace are stored.
	S3Location *string `field:"required" json:"s3Location" yaml:"s3Location"`
	// The ID of the workspace.
	WorkspaceId *string `field:"required" json:"workspaceId" yaml:"workspaceId"`
	// The description of the workspace.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Metadata that you can use to manage the workspace.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
}

