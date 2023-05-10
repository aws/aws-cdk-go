package awsiottwinmaker


// Properties for defining a `CfnScene`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnSceneProps := &CfnSceneProps{
//   	ContentLocation: jsii.String("contentLocation"),
//   	SceneId: jsii.String("sceneId"),
//   	WorkspaceId: jsii.String("workspaceId"),
//
//   	// the properties below are optional
//   	Capabilities: []*string{
//   		jsii.String("capabilities"),
//   	},
//   	Description: jsii.String("description"),
//   	SceneMetadata: map[string]*string{
//   		"sceneMetadataKey": jsii.String("sceneMetadata"),
//   	},
//   	Tags: map[string]*string{
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
	// `AWS::IoTTwinMaker::Scene.SceneMetadata`.
	SceneMetadata interface{} `field:"optional" json:"sceneMetadata" yaml:"sceneMetadata"`
	// The ComponentType tags.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
}

