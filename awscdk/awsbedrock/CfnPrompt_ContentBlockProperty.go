package awsbedrock


// A block of content for a message that you pass to, or receive from, a model with the [Converse](https://docs.aws.amazon.com/bedrock/latest/APIReference/API_runtime_Converse.html) or [ConverseStream](https://docs.aws.amazon.com/bedrock/latest/APIReference/API_runtime_ConverseStream.html) API operations.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   contentBlockProperty := &ContentBlockProperty{
//   	Text: jsii.String("text"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-prompt-contentblock.html
//
type CfnPrompt_ContentBlockProperty struct {
	// Text to include in the message.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-prompt-contentblock.html#cfn-bedrock-prompt-contentblock-text
	//
	Text *string `field:"required" json:"text" yaml:"text"`
}

