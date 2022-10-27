package awssns

import (
	"github.com/aws/aws-cdk-go/awscdk"
)

// Properties for defining a `CfnTopic`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var dataProtectionPolicy interface{}
//
//   cfnTopicProps := &cfnTopicProps{
//   	contentBasedDeduplication: jsii.Boolean(false),
//   	dataProtectionPolicy: dataProtectionPolicy,
//   	displayName: jsii.String("displayName"),
//   	fifoTopic: jsii.Boolean(false),
//   	kmsMasterKeyId: jsii.String("kmsMasterKeyId"),
//   	signatureVersion: jsii.String("signatureVersion"),
//   	subscription: []interface{}{
//   		&subscriptionProperty{
//   			endpoint: jsii.String("endpoint"),
//   			protocol: jsii.String("protocol"),
//   		},
//   	},
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   	topicName: jsii.String("topicName"),
//   }
//
type CfnTopicProps struct {
	// Enables content-based deduplication for FIFO topics.
	//
	// - By default, `ContentBasedDeduplication` is set to `false` . If you create a FIFO topic and this attribute is `false` , you must specify a value for the `MessageDeduplicationId` parameter for the [Publish](https://docs.aws.amazon.com/sns/latest/api/API_Publish.html) action.
	// - When you set `ContentBasedDeduplication` to `true` , Amazon SNS uses a SHA-256 hash to generate the `MessageDeduplicationId` using the body of the message (but not the attributes of the message).
	//
	// (Optional) To override the generated value, you can specify a value for the the `MessageDeduplicationId` parameter for the `Publish` action.
	ContentBasedDeduplication interface{} `field:"optional" json:"contentBasedDeduplication" yaml:"contentBasedDeduplication"`
	// `AWS::SNS::Topic.DataProtectionPolicy`.
	DataProtectionPolicy interface{} `field:"optional" json:"dataProtectionPolicy" yaml:"dataProtectionPolicy"`
	// The display name to use for an Amazon SNS topic with SMS subscriptions.
	DisplayName *string `field:"optional" json:"displayName" yaml:"displayName"`
	// Set to true to create a FIFO topic.
	FifoTopic interface{} `field:"optional" json:"fifoTopic" yaml:"fifoTopic"`
	// The ID of an AWS managed customer master key (CMK) for Amazon SNS or a custom CMK.
	//
	// For more information, see [Key terms](https://docs.aws.amazon.com/sns/latest/dg/sns-server-side-encryption.html#sse-key-terms) . For more examples, see `[KeyId](https://docs.aws.amazon.com/kms/latest/APIReference/API_DescribeKey.html#API_DescribeKey_RequestParameters)` in the *AWS Key Management Service API Reference* .
	//
	// This property applies only to [server-side-encryption](https://docs.aws.amazon.com/sns/latest/dg/sns-server-side-encryption.html) .
	KmsMasterKeyId *string `field:"optional" json:"kmsMasterKeyId" yaml:"kmsMasterKeyId"`
	// `AWS::SNS::Topic.SignatureVersion`.
	SignatureVersion *string `field:"optional" json:"signatureVersion" yaml:"signatureVersion"`
	// The Amazon SNS subscriptions (endpoints) for this topic.
	Subscription interface{} `field:"optional" json:"subscription" yaml:"subscription"`
	// The list of tags to add to a new topic.
	//
	// > To be able to tag a topic on creation, you must have the `sns:CreateTopic` and `sns:TagResource` permissions.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
	// The name of the topic you want to create.
	//
	// Topic names must include only uppercase and lowercase ASCII letters, numbers, underscores, and hyphens, and must be between 1 and 256 characters long. FIFO topic names must end with `.fifo` .
	//
	// If you don't specify a name, AWS CloudFormation generates a unique physical ID and uses that ID for the topic name. For more information, see [Name type](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-name.html) .
	//
	// > If you specify a name, you can't perform updates that require replacement of this resource. You can perform updates that require no or some interruption. If you must replace the resource, specify a new name.
	TopicName *string `field:"optional" json:"topicName" yaml:"topicName"`
}

