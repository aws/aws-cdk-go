package awswisdom

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnContent`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnContentProps := &CfnContentProps{
//   	KnowledgeBaseId: jsii.String("knowledgeBaseId"),
//   	Name: jsii.String("name"),
//
//   	// the properties below are optional
//   	Metadata: map[string]*string{
//   		"metadataKey": jsii.String("metadata"),
//   	},
//   	OverrideLinkOutUri: jsii.String("overrideLinkOutUri"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	Title: jsii.String("title"),
//   	UploadId: jsii.String("uploadId"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-wisdom-content.html
//
type CfnContentProps struct {
	// The identifier of the knowledge base.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-wisdom-content.html#cfn-wisdom-content-knowledgebaseid
	//
	KnowledgeBaseId *string `field:"required" json:"knowledgeBaseId" yaml:"knowledgeBaseId"`
	// The name of the content.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-wisdom-content.html#cfn-wisdom-content-name
	//
	Name *string `field:"required" json:"name" yaml:"name"`
	// A key/value map to store attributes without affecting tagging or recommendations.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-wisdom-content.html#cfn-wisdom-content-metadata
	//
	Metadata interface{} `field:"optional" json:"metadata" yaml:"metadata"`
	// The URI you want to use for the article.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-wisdom-content.html#cfn-wisdom-content-overridelinkouturi
	//
	OverrideLinkOutUri *string `field:"optional" json:"overrideLinkOutUri" yaml:"overrideLinkOutUri"`
	// The tags used to organize, track, or control access for this resource.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-wisdom-content.html#cfn-wisdom-content-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
	// The title of the content.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-wisdom-content.html#cfn-wisdom-content-title
	//
	Title *string `field:"optional" json:"title" yaml:"title"`
	// A pointer to the uploaded asset.
	//
	// This value is returned by StartContentUpload.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-wisdom-content.html#cfn-wisdom-content-uploadid
	//
	UploadId *string `field:"optional" json:"uploadId" yaml:"uploadId"`
}

