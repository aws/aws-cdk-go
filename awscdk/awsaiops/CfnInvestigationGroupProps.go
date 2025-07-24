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
	// Specify either the name or the ARN of the investigation group that you want to view.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-aiops-investigationgroup.html#cfn-aiops-investigationgroup-name
	//
	Name *string `field:"required" json:"name" yaml:"name"`
	// Use this property to integrate CloudWatch investigations with chat applications.
	//
	// This property is an array. For the first string, specify the ARN of an Amazon SNS topic. For the array of strings, specify the ARNs of one or more chat applications configurations that you want to associate with that topic. For more information about these configuration ARNs, see [Getting started with Amazon Q in chat applications](https://docs.aws.amazon.com/chatbot/latest/adminguide/getting-started.html) and [Resource type defined by AWS Chatbot](https://docs.aws.amazon.com/service-authorization/latest/reference/list_awschatbot.html#awschatbot-resources-for-iam-policies) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-aiops-investigationgroup.html#cfn-aiops-investigationgroup-chatbotnotificationchannels
	//
	ChatbotNotificationChannels interface{} `field:"optional" json:"chatbotNotificationChannels" yaml:"chatbotNotificationChannels"`
	// Number of `sourceAccountId` values that have been configured for cross-account access.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-aiops-investigationgroup.html#cfn-aiops-investigationgroup-crossaccountconfigurations
	//
	CrossAccountConfigurations interface{} `field:"optional" json:"crossAccountConfigurations" yaml:"crossAccountConfigurations"`
	// Specifies the customer managed AWS KMS key that the investigation group uses to encrypt data, if there is one.
	//
	// If not, the investigation group uses an AWS key to encrypt the data.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-aiops-investigationgroup.html#cfn-aiops-investigationgroup-encryptionconfig
	//
	EncryptionConfig interface{} `field:"optional" json:"encryptionConfig" yaml:"encryptionConfig"`
	// Returns the IAM resource policy that is associated with the specified investigation group.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-aiops-investigationgroup.html#cfn-aiops-investigationgroup-investigationgrouppolicy
	//
	InvestigationGroupPolicy *string `field:"optional" json:"investigationGroupPolicy" yaml:"investigationGroupPolicy"`
	// Specify `true` to enable CloudWatch investigations to have access to change events that are recorded by CloudTrail.
	//
	// The default is `true` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-aiops-investigationgroup.html#cfn-aiops-investigationgroup-iscloudtraileventhistoryenabled
	//
	IsCloudTrailEventHistoryEnabled interface{} `field:"optional" json:"isCloudTrailEventHistoryEnabled" yaml:"isCloudTrailEventHistoryEnabled"`
	// Specifies how long that investigation data is kept.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-aiops-investigationgroup.html#cfn-aiops-investigationgroup-retentionindays
	//
	RetentionInDays *float64 `field:"optional" json:"retentionInDays" yaml:"retentionInDays"`
	// The ARN of the IAM role that the investigation group uses for permissions to gather data.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-aiops-investigationgroup.html#cfn-aiops-investigationgroup-rolearn
	//
	RoleArn *string `field:"optional" json:"roleArn" yaml:"roleArn"`
	// Displays the custom tag keys for custom applications in your system that you have specified in the investigation group.
	//
	// Resource tags help CloudWatch investigations narrow the search space when it is unable to discover definite relationships between resources.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-aiops-investigationgroup.html#cfn-aiops-investigationgroup-tagkeyboundaries
	//
	TagKeyBoundaries *[]*string `field:"optional" json:"tagKeyBoundaries" yaml:"tagKeyBoundaries"`
	// The list of key-value pairs to associate with the resource.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-aiops-investigationgroup.html#cfn-aiops-investigationgroup-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

