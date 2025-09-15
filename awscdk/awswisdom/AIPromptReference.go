package awswisdom


// A reference to a AIPrompt resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   aIPromptReference := &AIPromptReference{
//   	AiPromptArn: jsii.String("aiPromptArn"),
//   	AiPromptId: jsii.String("aiPromptId"),
//   	AssistantId: jsii.String("assistantId"),
//   }
//
type AIPromptReference struct {
	// The ARN of the AIPrompt resource.
	AiPromptArn *string `field:"required" json:"aiPromptArn" yaml:"aiPromptArn"`
	// The AIPromptId of the AIPrompt resource.
	AiPromptId *string `field:"required" json:"aiPromptId" yaml:"aiPromptId"`
	// The AssistantId of the AIPrompt resource.
	AssistantId *string `field:"required" json:"assistantId" yaml:"assistantId"`
}

