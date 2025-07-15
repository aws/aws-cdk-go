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
	// A name for the investigation group.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-aiops-investigationgroup.html#cfn-aiops-investigationgroup-name
	//
	Name *string `field:"required" json:"name" yaml:"name"`
	// Use this property to integrate Amazon Q Developer operational investigations with Amazon Q in chat applications.
	//
	// This property is an array. For the first string, specify the ARN of an Amazon SNS topic. For the array of strings, specify the ARNs of one or more Amazon Q in chat applications configurations that you want to associate with that topic. For more information about these configuration ARNs, see [Getting started with Amazon Q in chat applications](https://docs.aws.amazon.com/chatbot/latest/adminguide/getting-started.html) and [Resource type defined by AWS Chatbot](https://docs.aws.amazon.com/service-authorization/latest/reference/list_awschatbot.html#awschatbot-resources-for-iam-policies) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-aiops-investigationgroup.html#cfn-aiops-investigationgroup-chatbotnotificationchannels
	//
	ChatbotNotificationChannels interface{} `field:"optional" json:"chatbotNotificationChannels" yaml:"chatbotNotificationChannels"`
	// An array of cross account configurations.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-aiops-investigationgroup.html#cfn-aiops-investigationgroup-crossaccountconfigurations
	//
	CrossAccountConfigurations interface{} `field:"optional" json:"crossAccountConfigurations" yaml:"crossAccountConfigurations"`
	// Use this property to specify a customer managed AWS KMS key to encrypt your investigation data.
	//
	// If you omit this property, Amazon Q Developer operational investigations will use an AWS key to encrypt the data. For more information, see [Encryption of investigation data](https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/Investigations-Security.html#Investigations-KMS) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-aiops-investigationgroup.html#cfn-aiops-investigationgroup-encryptionconfig
	//
	EncryptionConfig interface{} `field:"optional" json:"encryptionConfig" yaml:"encryptionConfig"`
	// Investigation Group policy.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-aiops-investigationgroup.html#cfn-aiops-investigationgroup-investigationgrouppolicy
	//
	InvestigationGroupPolicy *string `field:"optional" json:"investigationGroupPolicy" yaml:"investigationGroupPolicy"`
	// Specify `true` to enable Amazon Q Developer operational investigations to have access to change events that are recorded by CloudTrail .
	//
	// The default is `true` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-aiops-investigationgroup.html#cfn-aiops-investigationgroup-iscloudtraileventhistoryenabled
	//
	IsCloudTrailEventHistoryEnabled interface{} `field:"optional" json:"isCloudTrailEventHistoryEnabled" yaml:"isCloudTrailEventHistoryEnabled"`
	// Specify how long that investigation data is kept. For more information, see [Operational investigation data retention](https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/Investigations-Retention.html) .
	//
	// If you omit this parameter, the default of 90 days is used.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-aiops-investigationgroup.html#cfn-aiops-investigationgroup-retentionindays
	//
	RetentionInDays *float64 `field:"optional" json:"retentionInDays" yaml:"retentionInDays"`
	// Specify the ARN of the IAM role that Amazon Q Developer operational investigations will use when it gathers investigation data.
	//
	// The permissions in this role determine which of your resources that Amazon Q Developer operational investigations will have access to during investigations.
	//
	// For more information, see [How to control what data Amazon Q has access to during investigations](https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/Investigations-Security.html#Investigations-Security-Data) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-aiops-investigationgroup.html#cfn-aiops-investigationgroup-rolearn
	//
	RoleArn *string `field:"optional" json:"roleArn" yaml:"roleArn"`
	// Enter the existing custom tag keys for custom applications in your system.
	//
	// Resource tags help Amazon Q narrow the search space when it is unable to discover definite relationships between resources. For example, to discover that an Amazon ECS service depends on an Amazon RDS database, Amazon Q can discover this relationship using data sources such as X-Ray and CloudWatch Application Signals. However, if you haven't deployed these features, Amazon Q will attempt to identify possible relationships. Tag boundaries can be used to narrow the resources that will be discovered by Amazon Q in these cases.
	//
	// You don't need to enter tags created by myApplications or AWS CloudFormation , because Amazon Q can automatically detect those tags.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-aiops-investigationgroup.html#cfn-aiops-investigationgroup-tagkeyboundaries
	//
	TagKeyBoundaries *[]*string `field:"optional" json:"tagKeyBoundaries" yaml:"tagKeyBoundaries"`
	// A list of key-value pairs to associate with the investigation group.
	//
	// You can associate as many as 50 tags with an investigation group.
	//
	// Tags can help you organize and categorize your resources.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-aiops-investigationgroup.html#cfn-aiops-investigationgroup-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}

