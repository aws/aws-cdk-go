package awswisdom


// A reference to a AIAgentVersion resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   aIAgentVersionReference := &AIAgentVersionReference{
//   	AiAgentId: jsii.String("aiAgentId"),
//   	AssistantId: jsii.String("assistantId"),
//   	VersionNumber: jsii.String("versionNumber"),
//   }
//
type AIAgentVersionReference struct {
	// The AIAgentId of the AIAgentVersion resource.
	AiAgentId *string `field:"required" json:"aiAgentId" yaml:"aiAgentId"`
	// The AssistantId of the AIAgentVersion resource.
	AssistantId *string `field:"required" json:"assistantId" yaml:"assistantId"`
	// The VersionNumber of the AIAgentVersion resource.
	VersionNumber *string `field:"required" json:"versionNumber" yaml:"versionNumber"`
}

