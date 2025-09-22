package awswisdom


// A reference to a Assistant resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   assistantReference := &AssistantReference{
//   	AssistantArn: jsii.String("assistantArn"),
//   	AssistantId: jsii.String("assistantId"),
//   }
//
type AssistantReference struct {
	// The ARN of the Assistant resource.
	AssistantArn *string `field:"required" json:"assistantArn" yaml:"assistantArn"`
	// The AssistantId of the Assistant resource.
	AssistantId *string `field:"required" json:"assistantId" yaml:"assistantId"`
}

