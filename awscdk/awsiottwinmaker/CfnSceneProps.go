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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iottwinmaker-scene.html
//
type CfnSceneProps struct {
	// The relative path that specifies the location of the content definition file.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iottwinmaker-scene.html#cfn-iottwinmaker-scene-contentlocation
	//
	ContentLocation *string `field:"required" json:"contentLocation" yaml:"contentLocation"`
	// The ID of the scene.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iottwinmaker-scene.html#cfn-iottwinmaker-scene-sceneid
	//
	SceneId *string `field:"required" json:"sceneId" yaml:"sceneId"`
	// The ID of the workspace.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iottwinmaker-scene.html#cfn-iottwinmaker-scene-workspaceid
	//
	WorkspaceId *string `field:"required" json:"workspaceId" yaml:"workspaceId"`
	// A list of capabilities that the scene uses to render.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iottwinmaker-scene.html#cfn-iottwinmaker-scene-capabilities
	//
	Capabilities *[]*string `field:"optional" json:"capabilities" yaml:"capabilities"`
	// The description of this scene.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iottwinmaker-scene.html#cfn-iottwinmaker-scene-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The scene metadata.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iottwinmaker-scene.html#cfn-iottwinmaker-scene-scenemetadata
	//
	SceneMetadata interface{} `field:"optional" json:"sceneMetadata" yaml:"sceneMetadata"`
	// The ComponentType tags.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iottwinmaker-scene.html#cfn-iottwinmaker-scene-tags
	//
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
}

