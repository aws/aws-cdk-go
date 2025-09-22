package awsiottwinmaker


// A reference to a Scene resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   sceneReference := &SceneReference{
//   	SceneArn: jsii.String("sceneArn"),
//   	SceneId: jsii.String("sceneId"),
//   	WorkspaceId: jsii.String("workspaceId"),
//   }
//
type SceneReference struct {
	// The ARN of the Scene resource.
	SceneArn *string `field:"required" json:"sceneArn" yaml:"sceneArn"`
	// The SceneId of the Scene resource.
	SceneId *string `field:"required" json:"sceneId" yaml:"sceneId"`
	// The WorkspaceId of the Scene resource.
	WorkspaceId *string `field:"required" json:"workspaceId" yaml:"workspaceId"`
}

