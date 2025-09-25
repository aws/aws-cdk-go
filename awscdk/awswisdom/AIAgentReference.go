package awswisdom


// A reference to a AIAgent resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   aIAgentReference := &AIAgentReference{
//   	AiAgentArn: jsii.String("aiAgentArn"),
//   	AiAgentId: jsii.String("aiAgentId"),
//   	AssistantId: jsii.String("assistantId"),
//   }
//
type AIAgentReference struct {
	// The ARN of the AIAgent resource.
	AiAgentArn *string `field:"required" json:"aiAgentArn" yaml:"aiAgentArn"`
	// The AIAgentId of the AIAgent resource.
	AiAgentId *string `field:"required" json:"aiAgentId" yaml:"aiAgentId"`
	// The AssistantId of the AIAgent resource.
	AssistantId *string `field:"required" json:"assistantId" yaml:"assistantId"`
}

