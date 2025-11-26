package previewawsbedrockmixins


// A message input, or returned from, a call to [Converse](https://docs.aws.amazon.com/bedrock/latest/APIReference/API_runtime_Converse.html) or [ConverseStream](https://docs.aws.amazon.com/bedrock/latest/APIReference/API_runtime_ConverseStream.html) .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   messageProperty := &MessageProperty{
//   	Content: []interface{}{
//   		&ContentBlockProperty{
//   			CachePoint: &CachePointBlockProperty{
//   				Type: jsii.String("type"),
//   			},
//   			Text: jsii.String("text"),
//   		},
//   	},
//   	Role: jsii.String("role"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-prompt-message.html
//
type CfnPromptPropsMixin_MessageProperty struct {
	// The message content. Note the following restrictions:.
	//
	// - You can include up to 20 images. Each image's size, height, and width must be no more than 3.75 MB, 8000 px, and 8000 px, respectively.
	// - You can include up to five documents. Each document's size must be no more than 4.5 MB.
	// - If you include a `ContentBlock` with a `document` field in the array, you must also include a `ContentBlock` with a `text` field.
	// - You can only include images and documents if the `role` is `user` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-prompt-message.html#cfn-bedrock-prompt-message-content
	//
	Content interface{} `field:"optional" json:"content" yaml:"content"`
	// The role that the message plays in the message.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-prompt-message.html#cfn-bedrock-prompt-message-role
	//
	Role *string `field:"optional" json:"role" yaml:"role"`
}

