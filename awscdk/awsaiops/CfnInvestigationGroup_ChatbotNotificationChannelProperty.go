package awsaiops


// Use this structure to integrate CloudWatch investigations with chat applications.
//
// This structure is a string array. For the first string, specify the ARN of an Amazon SNS topic. For the array of strings, specify the ARNs of one or more chat applications configurations that you want to associate with that topic. For more information about these configuration ARNs, see [Getting started with Amazon Q in chat applications](https://docs.aws.amazon.com/chatbot/latest/adminguide/getting-started.html) and [Resource type defined by AWS Chatbot](https://docs.aws.amazon.com/service-authorization/latest/reference/list_awschatbot.html#awschatbot-resources-for-iam-policies) .
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
	// Returns the Amazon Resource Name (ARN) of any third-party chat integrations configured for the account.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-aiops-investigationgroup-chatbotnotificationchannel.html#cfn-aiops-investigationgroup-chatbotnotificationchannel-chatconfigurationarns
	//
	ChatConfigurationArns *[]*string `field:"optional" json:"chatConfigurationArns" yaml:"chatConfigurationArns"`
	// Returns the ARN of an Amazon  topic used for third-party chat integrations.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-aiops-investigationgroup-chatbotnotificationchannel.html#cfn-aiops-investigationgroup-chatbotnotificationchannel-snstopicarn
	//
	SnsTopicArn *string `field:"optional" json:"snsTopicArn" yaml:"snsTopicArn"`
}

