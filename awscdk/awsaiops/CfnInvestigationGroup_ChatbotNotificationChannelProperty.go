package awsaiops


// This structure is a string array.
//
// The first string is the ARN of a Amazon SNS topic. The array of strings display the ARNs of Amazon Q in chat applications configurations that are associated with that topic. For more information about these configuration ARNs, see [Getting started with Amazon Q in chat applications](https://docs.aws.amazon.com/chatbot/latest/adminguide/getting-started.html) and [Resource type defined by AWS Chatbot](https://docs.aws.amazon.com/service-authorization/latest/reference/list_awschatbot.html#awschatbot-resources-for-iam-policies) .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   chatbotNotificationChannelProperty := &ChatbotNotificationChannelProperty{
//   	ChatConfigurationArns: []*string{
//   		jsii.String("chatConfigurationArns"),
//   	},
//   	SnsTopicArn: jsii.String("snsTopicArn"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-aiops-investigationgroup-chatbotnotificationchannel.html
//
type CfnInvestigationGroup_ChatbotNotificationChannelProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-aiops-investigationgroup-chatbotnotificationchannel.html#cfn-aiops-investigationgroup-chatbotnotificationchannel-chatconfigurationarns
	//
	ChatConfigurationArns *[]*string `field:"optional" json:"chatConfigurationArns" yaml:"chatConfigurationArns"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-aiops-investigationgroup-chatbotnotificationchannel.html#cfn-aiops-investigationgroup-chatbotnotificationchannel-snstopicarn
	//
	SnsTopicArn *string `field:"optional" json:"snsTopicArn" yaml:"snsTopicArn"`
}

