package mixinsawswisdom

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for CfnAssistantAssociationPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   cfnAssistantAssociationMixinProps := &CfnAssistantAssociationMixinProps{
//   	AssistantId: jsii.String("assistantId"),
//   	Association: &AssociationDataProperty{
//   		KnowledgeBaseId: jsii.String("knowledgeBaseId"),
//   	},
//   	AssociationType: jsii.String("associationType"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-wisdom-assistantassociation.html
//
type CfnAssistantAssociationMixinProps struct {
	// The identifier of the Wisdom assistant.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-wisdom-assistantassociation.html#cfn-wisdom-assistantassociation-assistantid
	//
	AssistantId *string `field:"optional" json:"assistantId" yaml:"assistantId"`
	// The identifier of the associated resource.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-wisdom-assistantassociation.html#cfn-wisdom-assistantassociation-association
	//
	Association interface{} `field:"optional" json:"association" yaml:"association"`
	// The type of association.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-wisdom-assistantassociation.html#cfn-wisdom-assistantassociation-associationtype
	//
	AssociationType *string `field:"optional" json:"associationType" yaml:"associationType"`
	// The tags used to organize, track, or control access for this resource.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-wisdom-assistantassociation.html#cfn-wisdom-assistantassociation-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

