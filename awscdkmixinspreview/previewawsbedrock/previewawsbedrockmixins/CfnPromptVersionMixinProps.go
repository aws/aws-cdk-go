package previewawsbedrockmixins


// Properties for CfnPromptVersionPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnPromptVersionMixinProps := &CfnPromptVersionMixinProps{
//   	Description: jsii.String("description"),
//   	PromptArn: jsii.String("promptArn"),
//   	Tags: map[string]*string{
//   		"tagsKey": jsii.String("tags"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrock-promptversion.html
//
type CfnPromptVersionMixinProps struct {
	// The description of the prompt version.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrock-promptversion.html#cfn-bedrock-promptversion-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The Amazon Resource Name (ARN) of the version of the prompt.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrock-promptversion.html#cfn-bedrock-promptversion-promptarn
	//
	PromptArn *string `field:"optional" json:"promptArn" yaml:"promptArn"`
	// A map of tags attached to the prompt version and their values.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrock-promptversion.html#cfn-bedrock-promptversion-tags
	//
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
}

