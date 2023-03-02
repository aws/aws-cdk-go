package awsiottwinmaker


// Properties for defining a `CfnScene`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnSceneProps := &cfnSceneProps{
//   	contentLocation: jsii.String("contentLocation"),
//   	sceneId: jsii.String("sceneId"),
//   	workspaceId: jsii.String("workspaceId"),
//
//   	// the properties below are optional
//   	capabilities: []*string{
//   		jsii.String("capabilities"),
//   	},
//   	description: jsii.String("description"),
//   	tags: map[string]*string{
//   		"tagsKey": jsii.String("tags"),
//   	},
//   }
//
type CfnSceneProps struct {
	// The relative path that specifies the location of the content definition file.
	ContentLocation *string `field:"required" json:"contentLocation" yaml:"contentLocation"`
	// The scene ID.
	SceneId *string `field:"required" json:"sceneId" yaml:"sceneId"`
	// The ID of the workspace.
	WorkspaceId *string `field:"required" json:"workspaceId" yaml:"workspaceId"`
	// A list of capabilities that the scene uses to render.
	Capabilities *[]*string `field:"optional" json:"capabilities" yaml:"capabilities"`
	// The description of this scene.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The ComponentType tags.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
}

