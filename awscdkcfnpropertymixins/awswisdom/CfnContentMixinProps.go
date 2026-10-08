package awswisdom

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for CfnContentPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   cfnContentMixinProps := &CfnContentMixinProps{
//   	KnowledgeBaseId: jsii.String("knowledgeBaseId"),
//   	Metadata: map[string]*string{
//   		"metadataKey": jsii.String("metadata"),
//   	},
//   	Name: jsii.String("name"),
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
type CfnContentMixinProps struct {
	// The identifier of the knowledge base.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-wisdom-content.html#cfn-wisdom-content-knowledgebaseid
	//
	KnowledgeBaseId *string `field:"optional" json:"knowledgeBaseId" yaml:"knowledgeBaseId"`
	// A key/value map to store attributes without affecting tagging or recommendations.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-wisdom-content.html#cfn-wisdom-content-metadata
	//
	Metadata interface{} `field:"optional" json:"metadata" yaml:"metadata"`
	// The name of the content.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-wisdom-content.html#cfn-wisdom-content-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
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

