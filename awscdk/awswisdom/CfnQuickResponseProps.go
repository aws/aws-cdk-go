package awswisdom

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnQuickResponse`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnQuickResponseProps := &CfnQuickResponseProps{
//   	Content: &QuickResponseContentProviderProperty{
//   		Content: jsii.String("content"),
//   	},
//   	KnowledgeBaseArn: jsii.String("knowledgeBaseArn"),
//   	Name: jsii.String("name"),
//
//   	// the properties below are optional
//   	Channels: []*string{
//   		jsii.String("channels"),
//   	},
//   	ContentType: jsii.String("contentType"),
//   	Description: jsii.String("description"),
//   	GroupingConfiguration: &GroupingConfigurationProperty{
//   		Criteria: jsii.String("criteria"),
//   		Values: []*string{
//   			jsii.String("values"),
//   		},
//   	},
//   	IsActive: jsii.Boolean(false),
//   	Language: jsii.String("language"),
//   	ShortcutKey: jsii.String("shortcutKey"),
//   	Tags: []cfnTag{
//   		&cfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-wisdom-quickresponse.html
//
type CfnQuickResponseProps struct {
	// The content of the quick response.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-wisdom-quickresponse.html#cfn-wisdom-quickresponse-content
	//
	Content interface{} `field:"required" json:"content" yaml:"content"`
	// The Amazon Resource Name (ARN) of the knowledge base.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-wisdom-quickresponse.html#cfn-wisdom-quickresponse-knowledgebasearn
	//
	KnowledgeBaseArn *string `field:"required" json:"knowledgeBaseArn" yaml:"knowledgeBaseArn"`
	// The name of the quick response.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-wisdom-quickresponse.html#cfn-wisdom-quickresponse-name
	//
	Name *string `field:"required" json:"name" yaml:"name"`
	// The Amazon Connect contact channels this quick response applies to.
	//
	// The supported contact channel types include `Chat` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-wisdom-quickresponse.html#cfn-wisdom-quickresponse-channels
	//
	Channels *[]*string `field:"optional" json:"channels" yaml:"channels"`
	// The media type of the quick response content.
	//
	// - Use `application/x.quickresponse;format=plain` for quick response written in plain text.
	// - Use `application/x.quickresponse;format=markdown` for quick response written in richtext.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-wisdom-quickresponse.html#cfn-wisdom-quickresponse-contenttype
	//
	ContentType *string `field:"optional" json:"contentType" yaml:"contentType"`
	// The description of the quick response.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-wisdom-quickresponse.html#cfn-wisdom-quickresponse-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The configuration information of the user groups that the quick response is accessible to.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-wisdom-quickresponse.html#cfn-wisdom-quickresponse-groupingconfiguration
	//
	GroupingConfiguration interface{} `field:"optional" json:"groupingConfiguration" yaml:"groupingConfiguration"`
	// Whether the quick response is active.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-wisdom-quickresponse.html#cfn-wisdom-quickresponse-isactive
	//
	IsActive interface{} `field:"optional" json:"isActive" yaml:"isActive"`
	// The language code value for the language in which the quick response is written.
	//
	// The supported language codes include `de_DE` , `en_US` , `es_ES` , `fr_FR` , `id_ID` , `it_IT` , `ja_JP` , `ko_KR` , `pt_BR` , `zh_CN` , `zh_TW`.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-wisdom-quickresponse.html#cfn-wisdom-quickresponse-language
	//
	Language *string `field:"optional" json:"language" yaml:"language"`
	// The shortcut key of the quick response.
	//
	// The value should be unique across the knowledge base.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-wisdom-quickresponse.html#cfn-wisdom-quickresponse-shortcutkey
	//
	ShortcutKey *string `field:"optional" json:"shortcutKey" yaml:"shortcutKey"`
	// The tags used to organize, track, or control access for this resource.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-wisdom-quickresponse.html#cfn-wisdom-quickresponse-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

