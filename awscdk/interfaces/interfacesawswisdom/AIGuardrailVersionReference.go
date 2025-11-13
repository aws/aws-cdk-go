package interfacesawswisdom


// A reference to a AIGuardrailVersion resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   aIGuardrailVersionReference := &AIGuardrailVersionReference{
//   	AiGuardrailId: jsii.String("aiGuardrailId"),
//   	AssistantId: jsii.String("assistantId"),
//   	VersionNumber: jsii.String("versionNumber"),
//   }
//
type AIGuardrailVersionReference struct {
	// The AIGuardrailId of the AIGuardrailVersion resource.
	AiGuardrailId *string `field:"required" json:"aiGuardrailId" yaml:"aiGuardrailId"`
	// The AssistantId of the AIGuardrailVersion resource.
	AssistantId *string `field:"required" json:"assistantId" yaml:"assistantId"`
	// The VersionNumber of the AIGuardrailVersion resource.
	VersionNumber *string `field:"required" json:"versionNumber" yaml:"versionNumber"`
}

