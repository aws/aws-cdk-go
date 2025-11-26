package previewawsbedrockmixins


// Contains configurations for instructions to provide the model for how to handle input.
//
// To learn more, see [Using the Converse API](https://docs.aws.amazon.com/bedrock/latest/userguide/conversation-inference-call.html) .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   systemContentBlockProperty := &SystemContentBlockProperty{
//   	CachePoint: &CachePointBlockProperty{
//   		Type: jsii.String("type"),
//   	},
//   	Text: jsii.String("text"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-promptversion-systemcontentblock.html
//
type CfnPromptVersionPropsMixin_SystemContentBlockProperty struct {
	// CachePoint to include in the system prompt.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-promptversion-systemcontentblock.html#cfn-bedrock-promptversion-systemcontentblock-cachepoint
	//
	CachePoint interface{} `field:"optional" json:"cachePoint" yaml:"cachePoint"`
	// A system prompt for the model.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-promptversion-systemcontentblock.html#cfn-bedrock-promptversion-systemcontentblock-text
	//
	Text *string `field:"optional" json:"text" yaml:"text"`
}

