package awswisdom


// A reference to a AssistantAssociation resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   assistantAssociationReference := &AssistantAssociationReference{
//   	AssistantAssociationArn: jsii.String("assistantAssociationArn"),
//   	AssistantAssociationId: jsii.String("assistantAssociationId"),
//   	AssistantId: jsii.String("assistantId"),
//   }
//
type AssistantAssociationReference struct {
	// The ARN of the AssistantAssociation resource.
	AssistantAssociationArn *string `field:"required" json:"assistantAssociationArn" yaml:"assistantAssociationArn"`
	// The AssistantAssociationId of the AssistantAssociation resource.
	AssistantAssociationId *string `field:"required" json:"assistantAssociationId" yaml:"assistantAssociationId"`
	// The AssistantId of the AssistantAssociation resource.
	AssistantId *string `field:"required" json:"assistantId" yaml:"assistantId"`
}

