package awscases

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnRelatedItem`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnRelatedItemProps := &CfnRelatedItemProps{
//   	CaseId: jsii.String("caseId"),
//   	Content: &RelatedItemContentProperty{
//   		Comment: &CommentContentProperty{
//   			Body: jsii.String("body"),
//   			ContentType: jsii.String("contentType"),
//   		},
//   	},
//   	DomainId: jsii.String("domainId"),
//   	Type: jsii.String("type"),
//
//   	// the properties below are optional
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cases-relateditem.html
//
type CfnRelatedItemProps struct {
	// A unique identifier of the case.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cases-relateditem.html#cfn-cases-relateditem-caseid
	//
	CaseId *string `field:"required" json:"caseId" yaml:"caseId"`
	// Represents the content of a related item.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cases-relateditem.html#cfn-cases-relateditem-content
	//
	Content interface{} `field:"required" json:"content" yaml:"content"`
	// The unique identifier of the Cases domain.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cases-relateditem.html#cfn-cases-relateditem-domainid
	//
	DomainId *string `field:"required" json:"domainId" yaml:"domainId"`
	// The type of a related item.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cases-relateditem.html#cfn-cases-relateditem-type
	//
	Type *string `field:"required" json:"type" yaml:"type"`
	// A list of tags on the related item.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cases-relateditem.html#cfn-cases-relateditem-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

