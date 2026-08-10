package awsquicksight

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for CfnKnowledgeBasePropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   var template interface{}
//
//   cfnKnowledgeBaseMixinProps := &CfnKnowledgeBaseMixinProps{
//   	AccessControlConfiguration: &AccessControlConfigurationProperty{
//   		IsAclEnabled: jsii.Boolean(false),
//   	},
//   	AwsAccountId: jsii.String("awsAccountId"),
//   	DataSourceArn: jsii.String("dataSourceArn"),
//   	Description: jsii.String("description"),
//   	IsEmailNotificationOptedForIngestionFailures: jsii.Boolean(false),
//   	KnowledgeBaseConfiguration: &KnowledgeBaseConfigurationProperty{
//   		TemplateConfiguration: &KbTemplateConfigurationProperty{
//   			Template: template,
//   		},
//   	},
//   	KnowledgeBaseId: jsii.String("knowledgeBaseId"),
//   	MediaExtractionConfiguration: &MediaExtractionConfigurationProperty{
//   		AudioExtractionConfiguration: &AudioExtractionConfigurationProperty{
//   			AudioExtractionStatus: jsii.String("audioExtractionStatus"),
//   		},
//   		ImageExtractionConfiguration: &ImageExtractionConfigurationProperty{
//   			ImageExtractionStatus: jsii.String("imageExtractionStatus"),
//   		},
//   		VideoExtractionConfiguration: &VideoExtractionConfigurationProperty{
//   			VideoExtractionStatus: jsii.String("videoExtractionStatus"),
//   			VideoExtractionType: jsii.String("videoExtractionType"),
//   		},
//   	},
//   	Name: jsii.String("name"),
//   	Permissions: []interface{}{
//   		&ResourcePermissionProperty{
//   			Actions: []*string{
//   				jsii.String("actions"),
//   			},
//   			Principal: jsii.String("principal"),
//   			Resource: jsii.String("resource"),
//   		},
//   	},
//   	PrimaryOwnerArn: jsii.String("primaryOwnerArn"),
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-quicksight-knowledgebase.html
//
type CfnKnowledgeBaseMixinProps struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-quicksight-knowledgebase.html#cfn-quicksight-knowledgebase-accesscontrolconfiguration
	//
	AccessControlConfiguration interface{} `field:"optional" json:"accessControlConfiguration" yaml:"accessControlConfiguration"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-quicksight-knowledgebase.html#cfn-quicksight-knowledgebase-awsaccountid
	//
	AwsAccountId *string `field:"optional" json:"awsAccountId" yaml:"awsAccountId"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-quicksight-knowledgebase.html#cfn-quicksight-knowledgebase-datasourcearn
	//
	DataSourceArn *string `field:"optional" json:"dataSourceArn" yaml:"dataSourceArn"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-quicksight-knowledgebase.html#cfn-quicksight-knowledgebase-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-quicksight-knowledgebase.html#cfn-quicksight-knowledgebase-isemailnotificationoptedforingestionfailures
	//
	IsEmailNotificationOptedForIngestionFailures interface{} `field:"optional" json:"isEmailNotificationOptedForIngestionFailures" yaml:"isEmailNotificationOptedForIngestionFailures"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-quicksight-knowledgebase.html#cfn-quicksight-knowledgebase-knowledgebaseconfiguration
	//
	KnowledgeBaseConfiguration interface{} `field:"optional" json:"knowledgeBaseConfiguration" yaml:"knowledgeBaseConfiguration"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-quicksight-knowledgebase.html#cfn-quicksight-knowledgebase-knowledgebaseid
	//
	KnowledgeBaseId *string `field:"optional" json:"knowledgeBaseId" yaml:"knowledgeBaseId"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-quicksight-knowledgebase.html#cfn-quicksight-knowledgebase-mediaextractionconfiguration
	//
	MediaExtractionConfiguration interface{} `field:"optional" json:"mediaExtractionConfiguration" yaml:"mediaExtractionConfiguration"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-quicksight-knowledgebase.html#cfn-quicksight-knowledgebase-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-quicksight-knowledgebase.html#cfn-quicksight-knowledgebase-permissions
	//
	Permissions interface{} `field:"optional" json:"permissions" yaml:"permissions"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-quicksight-knowledgebase.html#cfn-quicksight-knowledgebase-primaryownerarn
	//
	PrimaryOwnerArn *string `field:"optional" json:"primaryOwnerArn" yaml:"primaryOwnerArn"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-quicksight-knowledgebase.html#cfn-quicksight-knowledgebase-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

