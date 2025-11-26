package previewawsiottwinmakermixins


// Properties for CfnScenePropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnSceneMixinProps := &CfnSceneMixinProps{
//   	Capabilities: []*string{
//   		jsii.String("capabilities"),
//   	},
//   	ContentLocation: jsii.String("contentLocation"),
//   	Description: jsii.String("description"),
//   	SceneId: jsii.String("sceneId"),
//   	SceneMetadata: map[string]*string{
//   		"sceneMetadataKey": jsii.String("sceneMetadata"),
//   	},
//   	Tags: map[string]*string{
//   		"tagsKey": jsii.String("tags"),
//   	},
//   	WorkspaceId: jsii.String("workspaceId"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iottwinmaker-scene.html
//
type CfnSceneMixinProps struct {
	// A list of capabilities that the scene uses to render.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iottwinmaker-scene.html#cfn-iottwinmaker-scene-capabilities
	//
	Capabilities *[]*string `field:"optional" json:"capabilities" yaml:"capabilities"`
	// The relative path that specifies the location of the content definition file.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iottwinmaker-scene.html#cfn-iottwinmaker-scene-contentlocation
	//
	ContentLocation *string `field:"optional" json:"contentLocation" yaml:"contentLocation"`
	// The description of this scene.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iottwinmaker-scene.html#cfn-iottwinmaker-scene-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The ID of the scene.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iottwinmaker-scene.html#cfn-iottwinmaker-scene-sceneid
	//
	SceneId *string `field:"optional" json:"sceneId" yaml:"sceneId"`
	// The scene metadata.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iottwinmaker-scene.html#cfn-iottwinmaker-scene-scenemetadata
	//
	SceneMetadata interface{} `field:"optional" json:"sceneMetadata" yaml:"sceneMetadata"`
	// The ComponentType tags.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iottwinmaker-scene.html#cfn-iottwinmaker-scene-tags
	//
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// The ID of the workspace.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iottwinmaker-scene.html#cfn-iottwinmaker-scene-workspaceid
	//
	WorkspaceId *string `field:"optional" json:"workspaceId" yaml:"workspaceId"`
}

