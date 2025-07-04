package awsaiops

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnInvestigationGroup`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnInvestigationGroupProps := &CfnInvestigationGroupProps{
//   	Name: jsii.String("name"),
//
//   	// the properties below are optional
//   	ChatbotNotificationChannels: []interface{}{
//   		&ChatbotNotificationChannelProperty{
//   			ChatConfigurationArns: []*string{
//   				jsii.String("chatConfigurationArns"),
//   			},
//   			SnsTopicArn: jsii.String("snsTopicArn"),
//   		},
//   	},
//   	CrossAccountConfigurations: []interface{}{
//   		&CrossAccountConfigurationProperty{
//   			SourceRoleArn: jsii.String("sourceRoleArn"),
//   		},
//   	},
//   	EncryptionConfig: &EncryptionConfigMapProperty{
//   		EncryptionConfigurationType: jsii.String("encryptionConfigurationType"),
//   		KmsKeyId: jsii.String("kmsKeyId"),
//   	},
//   	InvestigationGroupPolicy: jsii.String("investigationGroupPolicy"),
//   	IsCloudTrailEventHistoryEnabled: jsii.Boolean(false),
//   	RetentionInDays: jsii.Number(123),
//   	RoleArn: jsii.String("roleArn"),
//   	TagKeyBoundaries: []*string{
//   		jsii.String("tagKeyBoundaries"),
//   	},
//   	Tags: []cfnTag{
//   		&cfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-aiops-investigationgroup.html
//
type CfnInvestigationGroupProps struct {
	// User friendly name for resources.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-aiops-investigationgroup.html#cfn-aiops-investigationgroup-name
	//
	Name *string `field:"required" json:"name" yaml:"name"`
	// An array of key-value pairs of notification channels to apply to this resource.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-aiops-investigationgroup.html#cfn-aiops-investigationgroup-chatbotnotificationchannels
	//
	ChatbotNotificationChannels interface{} `field:"optional" json:"chatbotNotificationChannels" yaml:"chatbotNotificationChannels"`
	// An array of cross account configurations.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-aiops-investigationgroup.html#cfn-aiops-investigationgroup-crossaccountconfigurations
	//
	CrossAccountConfigurations interface{} `field:"optional" json:"crossAccountConfigurations" yaml:"crossAccountConfigurations"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-aiops-investigationgroup.html#cfn-aiops-investigationgroup-encryptionconfig
	//
	EncryptionConfig interface{} `field:"optional" json:"encryptionConfig" yaml:"encryptionConfig"`
	// Investigation Group policy.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-aiops-investigationgroup.html#cfn-aiops-investigationgroup-investigationgrouppolicy
	//
	InvestigationGroupPolicy *string `field:"optional" json:"investigationGroupPolicy" yaml:"investigationGroupPolicy"`
	// Flag to enable cloud trail history.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-aiops-investigationgroup.html#cfn-aiops-investigationgroup-iscloudtraileventhistoryenabled
	//
	IsCloudTrailEventHistoryEnabled interface{} `field:"optional" json:"isCloudTrailEventHistoryEnabled" yaml:"isCloudTrailEventHistoryEnabled"`
	// The number of days to retain the investigation group.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-aiops-investigationgroup.html#cfn-aiops-investigationgroup-retentionindays
	//
	RetentionInDays *float64 `field:"optional" json:"retentionInDays" yaml:"retentionInDays"`
	// The Investigation Role's ARN.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-aiops-investigationgroup.html#cfn-aiops-investigationgroup-rolearn
	//
	RoleArn *string `field:"optional" json:"roleArn" yaml:"roleArn"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-aiops-investigationgroup.html#cfn-aiops-investigationgroup-tagkeyboundaries
	//
	TagKeyBoundaries *[]*string `field:"optional" json:"tagKeyBoundaries" yaml:"tagKeyBoundaries"`
	// An array of key-value pairs to apply to this resource.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-aiops-investigationgroup.html#cfn-aiops-investigationgroup-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

