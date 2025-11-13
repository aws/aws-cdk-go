package interfacesawswisdom


// A reference to a AIPromptVersion resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   aIPromptVersionReference := &AIPromptVersionReference{
//   	AiPromptId: jsii.String("aiPromptId"),
//   	AssistantId: jsii.String("assistantId"),
//   	VersionNumber: jsii.String("versionNumber"),
//   }
//
type AIPromptVersionReference struct {
	// The AIPromptId of the AIPromptVersion resource.
	AiPromptId *string `field:"required" json:"aiPromptId" yaml:"aiPromptId"`
	// The AssistantId of the AIPromptVersion resource.
	AssistantId *string `field:"required" json:"assistantId" yaml:"assistantId"`
	// The VersionNumber of the AIPromptVersion resource.
	VersionNumber *string `field:"required" json:"versionNumber" yaml:"versionNumber"`
}

