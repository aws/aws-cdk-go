package awssns


// Properties for defining a `CfnSubscription`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var deliveryPolicy interface{}
//   var filterPolicy interface{}
//   var redrivePolicy interface{}
//   var replayPolicy interface{}
//
//   cfnSubscriptionProps := &CfnSubscriptionProps{
//   	Protocol: jsii.String("protocol"),
//   	TopicArn: jsii.String("topicArn"),
//
//   	// the properties below are optional
//   	DeliveryPolicy: deliveryPolicy,
//   	Endpoint: jsii.String("endpoint"),
//   	FilterPolicy: filterPolicy,
//   	FilterPolicyScope: jsii.String("filterPolicyScope"),
//   	RawMessageDelivery: jsii.Boolean(false),
//   	RedrivePolicy: redrivePolicy,
//   	Region: jsii.String("region"),
//   	ReplayPolicy: replayPolicy,
//   	SubscriptionRoleArn: jsii.String("subscriptionRoleArn"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sns-subscription.html
//
type CfnSubscriptionProps struct {
	// The subscription's protocol.
	//
	// For more information, see the `Protocol` parameter of the `[Subscribe](https://docs.aws.amazon.com/sns/latest/api/API_Subscribe.html)` action in the *Amazon SNS API Reference* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sns-subscription.html#cfn-sns-subscription-protocol
	//
	Protocol *string `field:"required" json:"protocol" yaml:"protocol"`
	// The ARN of the topic to subscribe to.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sns-subscription.html#cfn-sns-subscription-topicarn
	//
	TopicArn *string `field:"required" json:"topicArn" yaml:"topicArn"`
	// The delivery policy JSON assigned to the subscription.
	//
	// Enables the subscriber to define the message delivery retry strategy in the case of an HTTP/S endpoint subscribed to the topic. For more information, see `[GetSubscriptionAttributes](https://docs.aws.amazon.com/sns/latest/api/API_GetSubscriptionAttributes.html)` in the *Amazon SNS API Reference* and [Message delivery retries](https://docs.aws.amazon.com/sns/latest/dg/sns-message-delivery-retries.html) in the *Amazon SNS Developer Guide* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sns-subscription.html#cfn-sns-subscription-deliverypolicy
	//
	DeliveryPolicy interface{} `field:"optional" json:"deliveryPolicy" yaml:"deliveryPolicy"`
	// The subscription's endpoint.
	//
	// The endpoint value depends on the protocol that you specify. For more information, see the `Endpoint` parameter of the `[Subscribe](https://docs.aws.amazon.com/sns/latest/api/API_Subscribe.html)` action in the *Amazon SNS API Reference* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sns-subscription.html#cfn-sns-subscription-endpoint
	//
	Endpoint *string `field:"optional" json:"endpoint" yaml:"endpoint"`
	// The filter policy JSON assigned to the subscription.
	//
	// Enables the subscriber to filter out unwanted messages. For more information, see `[GetSubscriptionAttributes](https://docs.aws.amazon.com/sns/latest/api/API_GetSubscriptionAttributes.html)` in the *Amazon SNS API Reference* and [Message filtering](https://docs.aws.amazon.com/sns/latest/dg/sns-message-filtering.html) in the *Amazon SNS Developer Guide* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sns-subscription.html#cfn-sns-subscription-filterpolicy
	//
	FilterPolicy interface{} `field:"optional" json:"filterPolicy" yaml:"filterPolicy"`
	// This attribute lets you choose the filtering scope by using one of the following string value types:.
	//
	// - `MessageAttributes` (default) - The filter is applied on the message attributes.
	// - `MessageBody` - The filter is applied on the message body.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sns-subscription.html#cfn-sns-subscription-filterpolicyscope
	//
	FilterPolicyScope *string `field:"optional" json:"filterPolicyScope" yaml:"filterPolicyScope"`
	// When set to `true` , enables raw message delivery.
	//
	// Raw messages don't contain any JSON formatting and can be sent to Amazon SQS and HTTP/S endpoints. For more information, see `[GetSubscriptionAttributes](https://docs.aws.amazon.com/sns/latest/api/API_GetSubscriptionAttributes.html)` in the *Amazon SNS API Reference* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sns-subscription.html#cfn-sns-subscription-rawmessagedelivery
	//
	RawMessageDelivery interface{} `field:"optional" json:"rawMessageDelivery" yaml:"rawMessageDelivery"`
	// When specified, sends undeliverable messages to the specified Amazon SQS dead-letter queue.
	//
	// Messages that can't be delivered due to client errors (for example, when the subscribed endpoint is unreachable) or server errors (for example, when the service that powers the subscribed endpoint becomes unavailable) are held in the dead-letter queue for further analysis or reprocessing.
	//
	// For more information about the redrive policy and dead-letter queues, see [Amazon SQS dead-letter queues](https://docs.aws.amazon.com/AWSSimpleQueueService/latest/SQSDeveloperGuide/sqs-dead-letter-queues.html) in the *Amazon SQS Developer Guide* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sns-subscription.html#cfn-sns-subscription-redrivepolicy
	//
	RedrivePolicy interface{} `field:"optional" json:"redrivePolicy" yaml:"redrivePolicy"`
	// For cross-region subscriptions, the region in which the topic resides.
	//
	// If no region is specified, AWS CloudFormation uses the region of the caller as the default.
	//
	// If you perform an update operation that only updates the `Region` property of a `AWS::SNS::Subscription` resource, that operation will fail unless you are either:
	//
	// - Updating the `Region` from `NULL` to the caller region.
	// - Updating the `Region` from the caller region to `NULL` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sns-subscription.html#cfn-sns-subscription-region
	//
	Region *string `field:"optional" json:"region" yaml:"region"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sns-subscription.html#cfn-sns-subscription-replaypolicy
	//
	ReplayPolicy interface{} `field:"optional" json:"replayPolicy" yaml:"replayPolicy"`
	// This property applies only to Amazon Data Firehose delivery stream subscriptions.
	//
	// Specify the ARN of the IAM role that has the following:
	//
	// - Permission to write to the Amazon Data Firehose delivery stream
	// - Amazon SNS listed as a trusted entity
	//
	// Specifying a valid ARN for this attribute is required for Firehose delivery stream subscriptions. For more information, see [Fanout to Amazon Data Firehose delivery streams](https://docs.aws.amazon.com/sns/latest/dg/sns-firehose-as-subscriber.html) in the *Amazon SNS Developer Guide.*
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sns-subscription.html#cfn-sns-subscription-subscriptionrolearn
	//
	SubscriptionRoleArn *string `field:"optional" json:"subscriptionRoleArn" yaml:"subscriptionRoleArn"`
}

