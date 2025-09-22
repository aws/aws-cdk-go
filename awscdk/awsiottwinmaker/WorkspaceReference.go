package awsiottwinmaker


// A reference to a Workspace resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   workspaceReference := &WorkspaceReference{
//   	WorkspaceArn: jsii.String("workspaceArn"),
//   	WorkspaceId: jsii.String("workspaceId"),
//   }
//
type WorkspaceReference struct {
	// The ARN of the Workspace resource.
	WorkspaceArn *string `field:"required" json:"workspaceArn" yaml:"workspaceArn"`
	// The WorkspaceId of the Workspace resource.
	WorkspaceId *string `field:"required" json:"workspaceId" yaml:"workspaceId"`
}

