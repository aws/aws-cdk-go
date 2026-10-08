package awswisdom

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for CfnContentAssociationPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   cfnContentAssociationMixinProps := &CfnContentAssociationMixinProps{
//   	Association: &ContentAssociationContentsProperty{
//   		AmazonConnectGuideAssociation: &AmazonConnectGuideAssociationDataProperty{
//   			FlowId: jsii.String("flowId"),
//   		},
//   	},
//   	AssociationType: jsii.String("associationType"),
//   	ContentId: jsii.String("contentId"),
//   	KnowledgeBaseId: jsii.String("knowledgeBaseId"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-wisdom-contentassociation.html
//
type CfnContentAssociationMixinProps struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-wisdom-contentassociation.html#cfn-wisdom-contentassociation-association
	//
	Association interface{} `field:"optional" json:"association" yaml:"association"`
	// The type of association.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-wisdom-contentassociation.html#cfn-wisdom-contentassociation-associationtype
	//
	AssociationType *string `field:"optional" json:"associationType" yaml:"associationType"`
	// The identifier of the content.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-wisdom-contentassociation.html#cfn-wisdom-contentassociation-contentid
	//
	ContentId *string `field:"optional" json:"contentId" yaml:"contentId"`
	// The identifier of the knowledge base.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-wisdom-contentassociation.html#cfn-wisdom-contentassociation-knowledgebaseid
	//
	KnowledgeBaseId *string `field:"optional" json:"knowledgeBaseId" yaml:"knowledgeBaseId"`
	// The tags used to organize, track, or control access for this resource.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-wisdom-contentassociation.html#cfn-wisdom-contentassociation-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

