package awsbedrock


// Chat prompt Message.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   messageProperty := &MessageProperty{
//   	Content: []interface{}{
//   		&ContentBlockProperty{
//   			Text: jsii.String("text"),
//   		},
//   	},
//   	Role: jsii.String("role"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-promptversion-message.html
//
type CfnPromptVersion_MessageProperty struct {
	// List of Content Blocks.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-promptversion-message.html#cfn-bedrock-promptversion-message-content
	//
	Content interface{} `field:"required" json:"content" yaml:"content"`
	// Conversation roles for the chat prompt.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-promptversion-message.html#cfn-bedrock-promptversion-message-role
	//
	Role *string `field:"required" json:"role" yaml:"role"`
}

