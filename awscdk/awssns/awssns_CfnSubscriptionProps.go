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
//
//   cfnSubscriptionProps := &cfnSubscriptionProps{
//   	protocol: jsii.String("protocol"),
//   	topicArn: jsii.String("topicArn"),
//
//   	// the properties below are optional
//   	deliveryPolicy: deliveryPolicy,
//   	endpoint: jsii.String("endpoint"),
//   	filterPolicy: filterPolicy,
//   	rawMessageDelivery: jsii.Boolean(false),
//   	redrivePolicy: redrivePolicy,
//   	region: jsii.String("region"),
//   	subscriptionRoleArn: jsii.String("subscriptionRoleArn"),
//   }
//
type CfnSubscriptionProps struct {
	// The subscription's protocol.
	//
	// For more information, see the `Protocol` parameter of the `[Subscribe](https://docs.aws.amazon.com/sns/latest/api/API_Subscribe.html)` action in the *Amazon SNS API Reference* .
	Protocol *string `field:"required" json:"protocol" yaml:"protocol"`
	// The ARN of the topic to subscribe to.
	TopicArn *string `field:"required" json:"topicArn" yaml:"topicArn"`
	// The delivery policy JSON assigned to the subscription.
	//
	// Enables the subscriber to define the message delivery retry strategy in the case of an HTTP/S endpoint subscribed to the topic. For more information, see `[GetSubscriptionAttributes](https://docs.aws.amazon.com/sns/latest/api/API_GetSubscriptionAttributes.html)` in the *Amazon SNS API Reference* and [Message delivery retries](https://docs.aws.amazon.com/sns/latest/dg/sns-message-delivery-retries.html) in the *Amazon SNS Developer Guide* .
	DeliveryPolicy interface{} `field:"optional" json:"deliveryPolicy" yaml:"deliveryPolicy"`
	// The subscription's endpoint.
	//
	// The endpoint value depends on the protocol that you specify. For more information, see the `Endpoint` parameter of the `[Subscribe](https://docs.aws.amazon.com/sns/latest/api/API_Subscribe.html)` action in the *Amazon SNS API Reference* .
	Endpoint *string `field:"optional" json:"endpoint" yaml:"endpoint"`
	// The filter policy JSON assigned to the subscription.
	//
	// Enables the subscriber to filter out unwanted messages. For more information, see `[GetSubscriptionAttributes](https://docs.aws.amazon.com/sns/latest/api/API_GetSubscriptionAttributes.html)` in the *Amazon SNS API Reference* and [Message filtering](https://docs.aws.amazon.com/sns/latest/dg/sns-message-filtering.html) in the *Amazon SNS Developer Guide* .
	FilterPolicy interface{} `field:"optional" json:"filterPolicy" yaml:"filterPolicy"`
	// When set to `true` , enables raw message delivery.
	//
	// Raw messages don't contain any JSON formatting and can be sent to Amazon SQS and HTTP/S endpoints. For more information, see `[GetSubscriptionAttributes](https://docs.aws.amazon.com/sns/latest/api/API_GetSubscriptionAttributes.html)` in the *Amazon SNS API Reference* .
	RawMessageDelivery interface{} `field:"optional" json:"rawMessageDelivery" yaml:"rawMessageDelivery"`
	// When specified, sends undeliverable messages to the specified Amazon SQS dead-letter queue.
	//
	// Messages that can't be delivered due to client errors (for example, when the subscribed endpoint is unreachable) or server errors (for example, when the service that powers the subscribed endpoint becomes unavailable) are held in the dead-letter queue for further analysis or reprocessing.
	//
	// For more information about the redrive policy and dead-letter queues, see [Amazon SQS dead-letter queues](https://docs.aws.amazon.com/AWSSimpleQueueService/latest/SQSDeveloperGuide/sqs-dead-letter-queues.html) in the *Amazon SQS Developer Guide* .
	RedrivePolicy interface{} `field:"optional" json:"redrivePolicy" yaml:"redrivePolicy"`
	// For cross-region subscriptions, the region in which the topic resides.
	//
	// If no region is specified, AWS CloudFormation uses the region of the caller as the default.
	//
	// If you perform an update operation that only updates the `Region` property of a `AWS::SNS::Subscription` resource, that operation will fail unless you are either:
	//
	// - Updating the `Region` from `NULL` to the caller region.
	// - Updating the `Region` from the caller region to `NULL` .
	Region *string `field:"optional" json:"region" yaml:"region"`
	// This property applies only to Amazon Kinesis Data Firehose delivery stream subscriptions.
	//
	// Specify the ARN of the IAM role that has the following:
	//
	// - Permission to write to the Amazon Kinesis Data Firehose delivery stream
	// - Amazon SNS listed as a trusted entity
	//
	// Specifying a valid ARN for this attribute is required for Kinesis Data Firehose delivery stream subscriptions. For more information, see [Fanout to Amazon Kinesis Data Firehose delivery streams](https://docs.aws.amazon.com/sns/latest/dg/sns-firehose-as-subscriber.html) in the *Amazon SNS Developer Guide.*
	SubscriptionRoleArn *string `field:"optional" json:"subscriptionRoleArn" yaml:"subscriptionRoleArn"`
}

