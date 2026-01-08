package awswisdom


// A union type that currently has a single argument, which is the knowledge base ID.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   associationDataProperty := &AssociationDataProperty{
//   	ExternalBedrockKnowledgeBaseConfig: &ExternalBedrockKnowledgeBaseConfigProperty{
//   		AccessRoleArn: jsii.String("accessRoleArn"),
//   		BedrockKnowledgeBaseArn: jsii.String("bedrockKnowledgeBaseArn"),
//   	},
//   	KnowledgeBaseId: jsii.String("knowledgeBaseId"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wisdom-assistantassociation-associationdata.html
//
type CfnAssistantAssociation_AssociationDataProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wisdom-assistantassociation-associationdata.html#cfn-wisdom-assistantassociation-associationdata-externalbedrockknowledgebaseconfig
	//
	ExternalBedrockKnowledgeBaseConfig interface{} `field:"optional" json:"externalBedrockKnowledgeBaseConfig" yaml:"externalBedrockKnowledgeBaseConfig"`
	// The identifier of the knowledge base.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wisdom-assistantassociation-associationdata.html#cfn-wisdom-assistantassociation-associationdata-knowledgebaseid
	//
	KnowledgeBaseId *string `field:"optional" json:"knowledgeBaseId" yaml:"knowledgeBaseId"`
}

