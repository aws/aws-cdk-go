package previewawsbedrockmixins


// Contains a key-value pair that defines a metadata tag and value to attach to a prompt variant.
//
// For more information, see [Create a prompt using Prompt management](https://docs.aws.amazon.com/bedrock/latest/userguide/prompt-management-create.html) .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   promptMetadataEntryProperty := &PromptMetadataEntryProperty{
//   	Key: jsii.String("key"),
//   	Value: jsii.String("value"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-promptversion-promptmetadataentry.html
//
type CfnPromptVersionPropsMixin_PromptMetadataEntryProperty struct {
	// The key of a metadata tag for a prompt variant.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-promptversion-promptmetadataentry.html#cfn-bedrock-promptversion-promptmetadataentry-key
	//
	Key *string `field:"optional" json:"key" yaml:"key"`
	// The value of a metadata tag for a prompt variant.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-promptversion-promptmetadataentry.html#cfn-bedrock-promptversion-promptmetadataentry-value
	//
	Value *string `field:"optional" json:"value" yaml:"value"`
}

