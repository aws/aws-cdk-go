package awssns

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnTopic`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var archivePolicy interface{}
//   var dataProtectionPolicy interface{}
//
//   cfnTopicProps := &CfnTopicProps{
//   	ArchivePolicy: archivePolicy,
//   	ContentBasedDeduplication: jsii.Boolean(false),
//   	DataProtectionPolicy: dataProtectionPolicy,
//   	DeliveryStatusLogging: []interface{}{
//   		&LoggingConfigProperty{
//   			Protocol: jsii.String("protocol"),
//
//   			// the properties below are optional
//   			FailureFeedbackRoleArn: jsii.String("failureFeedbackRoleArn"),
//   			SuccessFeedbackRoleArn: jsii.String("successFeedbackRoleArn"),
//   			SuccessFeedbackSampleRate: jsii.String("successFeedbackSampleRate"),
//   		},
//   	},
//   	DisplayName: jsii.String("displayName"),
//   	FifoTopic: jsii.Boolean(false),
//   	KmsMasterKeyId: jsii.String("kmsMasterKeyId"),
//   	SignatureVersion: jsii.String("signatureVersion"),
//   	Subscription: []interface{}{
//   		&SubscriptionProperty{
//   			Endpoint: jsii.String("endpoint"),
//   			Protocol: jsii.String("protocol"),
//   		},
//   	},
//   	Tags: []cfnTag{
//   		&cfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	TopicName: jsii.String("topicName"),
//   	TracingConfig: jsii.String("tracingConfig"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sns-topic.html
//
type CfnTopicProps struct {
	// The archive policy determines the number of days Amazon SNS retains messages.
	//
	// You can set a retention period from 1 to 365 days.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sns-topic.html#cfn-sns-topic-archivepolicy
	//
	ArchivePolicy interface{} `field:"optional" json:"archivePolicy" yaml:"archivePolicy"`
	// Enables content-based deduplication for FIFO topics.
	//
	// - By default, `ContentBasedDeduplication` is set to `false` . If you create a FIFO topic and this attribute is `false` , you must specify a value for the `MessageDeduplicationId` parameter for the [Publish](https://docs.aws.amazon.com/sns/latest/api/API_Publish.html) action.
	// - When you set `ContentBasedDeduplication` to `true` , Amazon SNS uses a SHA-256 hash to generate the `MessageDeduplicationId` using the body of the message (but not the attributes of the message).
	//
	// (Optional) To override the generated value, you can specify a value for the the `MessageDeduplicationId` parameter for the `Publish` action.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sns-topic.html#cfn-sns-topic-contentbaseddeduplication
	//
	ContentBasedDeduplication interface{} `field:"optional" json:"contentBasedDeduplication" yaml:"contentBasedDeduplication"`
	// The body of the policy document you want to use for this topic.
	//
	// You can only add one policy per topic.
	//
	// The policy must be in JSON string format.
	//
	// Length Constraints: Maximum length of 30,720.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sns-topic.html#cfn-sns-topic-dataprotectionpolicy
	//
	DataProtectionPolicy interface{} `field:"optional" json:"dataProtectionPolicy" yaml:"dataProtectionPolicy"`
	// The `DeliveryStatusLogging` configuration enables you to log the delivery status of messages sent from your Amazon SNS topic to subscribed endpoints with the following supported delivery protocols:.
	//
	// - HTTP
	// - Amazon Kinesis Data Firehose
	// - AWS Lambda
	// - Platform application endpoint
	// - Amazon Simple Queue Service
	//
	// Once configured, log entries are sent to Amazon CloudWatch Logs.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sns-topic.html#cfn-sns-topic-deliverystatuslogging
	//
	DeliveryStatusLogging interface{} `field:"optional" json:"deliveryStatusLogging" yaml:"deliveryStatusLogging"`
	// The display name to use for an Amazon SNS topic with SMS subscriptions.
	//
	// The display name must be maximum 100 characters long, including hyphens (-), underscores (_), spaces, and tabs.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sns-topic.html#cfn-sns-topic-displayname
	//
	DisplayName *string `field:"optional" json:"displayName" yaml:"displayName"`
	// Set to true to create a FIFO topic.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sns-topic.html#cfn-sns-topic-fifotopic
	//
	FifoTopic interface{} `field:"optional" json:"fifoTopic" yaml:"fifoTopic"`
	// The ID of an AWS managed customer master key (CMK) for Amazon SNS or a custom CMK.
	//
	// For more information, see [Key terms](https://docs.aws.amazon.com/sns/latest/dg/sns-server-side-encryption.html#sse-key-terms) . For more examples, see `[KeyId](https://docs.aws.amazon.com/kms/latest/APIReference/API_DescribeKey.html#API_DescribeKey_RequestParameters)` in the *AWS Key Management Service API Reference* .
	//
	// This property applies only to [server-side-encryption](https://docs.aws.amazon.com/sns/latest/dg/sns-server-side-encryption.html) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sns-topic.html#cfn-sns-topic-kmsmasterkeyid
	//
	KmsMasterKeyId *string `field:"optional" json:"kmsMasterKeyId" yaml:"kmsMasterKeyId"`
	// The signature version corresponds to the hashing algorithm used while creating the signature of the notifications, subscription confirmations, or unsubscribe confirmation messages sent by Amazon SNS.
	//
	// By default, `SignatureVersion` is set to `1` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sns-topic.html#cfn-sns-topic-signatureversion
	//
	SignatureVersion *string `field:"optional" json:"signatureVersion" yaml:"signatureVersion"`
	// The Amazon SNS subscriptions (endpoints) for this topic.
	//
	// > If you specify the `Subscription` property in the `AWS::SNS::Topic` resource and it creates an associated subscription resource, the associated subscription is not deleted when the `AWS::SNS::Topic` resource is deleted.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sns-topic.html#cfn-sns-topic-subscription
	//
	Subscription interface{} `field:"optional" json:"subscription" yaml:"subscription"`
	// The list of tags to add to a new topic.
	//
	// > To be able to tag a topic on creation, you must have the `sns:CreateTopic` and `sns:TagResource` permissions.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sns-topic.html#cfn-sns-topic-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
	// The name of the topic you want to create.
	//
	// Topic names must include only uppercase and lowercase ASCII letters, numbers, underscores, and hyphens, and must be between 1 and 256 characters long. FIFO topic names must end with `.fifo` .
	//
	// If you don't specify a name, AWS CloudFormation generates a unique physical ID and uses that ID for the topic name. For more information, see [Name type](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-name.html) .
	//
	// > If you specify a name, you can't perform updates that require replacement of this resource. You can perform updates that require no or some interruption. If you must replace the resource, specify a new name.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sns-topic.html#cfn-sns-topic-topicname
	//
	TopicName *string `field:"optional" json:"topicName" yaml:"topicName"`
	// Tracing mode of an Amazon SNS topic.
	//
	// By default `TracingConfig` is set to `PassThrough` , and the topic passes through the tracing header it receives from an Amazon SNS publisher to its subscriptions. If set to `Active` , Amazon SNS will vend X-Ray segment data to topic owner account if the sampled flag in the tracing header is true.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sns-topic.html#cfn-sns-topic-tracingconfig
	//
	TracingConfig *string `field:"optional" json:"tracingConfig" yaml:"tracingConfig"`
}

