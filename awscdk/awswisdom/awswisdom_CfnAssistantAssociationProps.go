package awswisdom

import (
	"github.com/aws/aws-cdk-go/awscdk"
)

// Properties for defining a `CfnAssistantAssociation`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnAssistantAssociationProps := &cfnAssistantAssociationProps{
//   	assistantId: jsii.String("assistantId"),
//   	association: &associationDataProperty{
//   		knowledgeBaseId: jsii.String("knowledgeBaseId"),
//   	},
//   	associationType: jsii.String("associationType"),
//
//   	// the properties below are optional
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnAssistantAssociationProps struct {
	// The identifier of the Wisdom assistant.
	AssistantId *string `field:"required" json:"assistantId" yaml:"assistantId"`
	// The identifier of the associated resource.
	Association interface{} `field:"required" json:"association" yaml:"association"`
	// The type of association.
	AssociationType *string `field:"required" json:"associationType" yaml:"associationType"`
	// The tags used to organize, track, or control access for this resource.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

