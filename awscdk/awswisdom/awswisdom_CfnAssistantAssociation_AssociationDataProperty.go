package awswisdom


// A union type that currently has a single argument, which is the knowledge base ID.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   associationDataProperty := &associationDataProperty{
//   	knowledgeBaseId: jsii.String("knowledgeBaseId"),
//   }
//
type CfnAssistantAssociation_AssociationDataProperty struct {
	// The identifier of the knowledge base.
	KnowledgeBaseId *string `field:"required" json:"knowledgeBaseId" yaml:"knowledgeBaseId"`
}

