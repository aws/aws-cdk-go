package interfacesawswisdom


// A reference to a AIGuardrail resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   aIGuardrailReference := &AIGuardrailReference{
//   	AiGuardrailArn: jsii.String("aiGuardrailArn"),
//   	AiGuardrailId: jsii.String("aiGuardrailId"),
//   	AssistantId: jsii.String("assistantId"),
//   }
//
type AIGuardrailReference struct {
	// The ARN of the AIGuardrail resource.
	AiGuardrailArn *string `field:"required" json:"aiGuardrailArn" yaml:"aiGuardrailArn"`
	// The AIGuardrailId of the AIGuardrail resource.
	AiGuardrailId *string `field:"required" json:"aiGuardrailId" yaml:"aiGuardrailId"`
	// The AssistantId of the AIGuardrail resource.
	AssistantId *string `field:"required" json:"assistantId" yaml:"assistantId"`
}

