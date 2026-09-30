package awscases

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for CfnRelatedItemPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   cfnRelatedItemMixinProps := &CfnRelatedItemMixinProps{
//   	CaseId: jsii.String("caseId"),
//   	Content: &RelatedItemContentProperty{
//   		Comment: &CommentContentProperty{
//   			Body: jsii.String("body"),
//   			ContentType: jsii.String("contentType"),
//   		},
//   	},
//   	DomainId: jsii.String("domainId"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	Type: jsii.String("type"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cases-relateditem.html
//
type CfnRelatedItemMixinProps struct {
	// A unique identifier of the case.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cases-relateditem.html#cfn-cases-relateditem-caseid
	//
	CaseId *string `field:"optional" json:"caseId" yaml:"caseId"`
	// Represents the content of a related item.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cases-relateditem.html#cfn-cases-relateditem-content
	//
	Content interface{} `field:"optional" json:"content" yaml:"content"`
	// The unique identifier of the Cases domain.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cases-relateditem.html#cfn-cases-relateditem-domainid
	//
	DomainId *string `field:"optional" json:"domainId" yaml:"domainId"`
	// A list of tags on the related item.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cases-relateditem.html#cfn-cases-relateditem-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
	// The type of a related item.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cases-relateditem.html#cfn-cases-relateditem-type
	//
	Type *string `field:"optional" json:"type" yaml:"type"`
}

