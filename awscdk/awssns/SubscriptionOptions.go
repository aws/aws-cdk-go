package awssns

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awssqs"
)

// Options for creating a new subscription.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var filterOrPolicy filterOrPolicy
//   var queue queue
//   var subscriptionFilter subscriptionFilter
//
//   subscriptionOptions := &SubscriptionOptions{
//   	Endpoint: jsii.String("endpoint"),
//   	Protocol: awscdk.Aws_sns.SubscriptionProtocol_HTTP,
//
//   	// the properties below are optional
//   	DeadLetterQueue: queue,
//   	DeliveryPolicy: &DeliveryPolicy{
//   		HealthyRetryPolicy: &HealthyRetryPolicy{
//   			BackoffFunction: awscdk.*Aws_sns.BackoffFunction_ARITHMETIC,
//   			MaxDelayTarget: cdk.Duration_Minutes(jsii.Number(30)),
//   			MinDelayTarget: cdk.Duration_*Minutes(jsii.Number(30)),
//   			NumMaxDelayRetries: jsii.Number(123),
//   			NumMinDelayRetries: jsii.Number(123),
//   			NumNoDelayRetries: jsii.Number(123),
//   			NumRetries: jsii.Number(123),
//   		},
//   		RequestPolicy: &RequestPolicy{
//   			HeaderContentType: jsii.String("headerContentType"),
//   		},
//   		ThrottlePolicy: &ThrottlePolicy{
//   			MaxReceivesPerSecond: jsii.Number(123),
//   		},
//   	},
//   	FilterPolicy: map[string]*subscriptionFilter{
//   		"filterPolicyKey": subscriptionFilter,
//   	},
//   	FilterPolicyWithMessageBody: map[string]*filterOrPolicy{
//   		"filterPolicyWithMessageBodyKey": filterOrPolicy,
//   	},
//   	RawMessageDelivery: jsii.Boolean(false),
//   	Region: jsii.String("region"),
//   	SubscriptionRoleArn: jsii.String("subscriptionRoleArn"),
//   }
//
type SubscriptionOptions struct {
	// The subscription endpoint.
	//
	// The meaning of this value depends on the value for 'protocol'.
	Endpoint *string `field:"required" json:"endpoint" yaml:"endpoint"`
	// What type of subscription to add.
	Protocol SubscriptionProtocol `field:"required" json:"protocol" yaml:"protocol"`
	// Queue to be used as dead letter queue.
	//
	// If not passed no dead letter queue is enabled.
	// Default: - No dead letter queue enabled.
	//
	DeadLetterQueue awssqs.IQueue `field:"optional" json:"deadLetterQueue" yaml:"deadLetterQueue"`
	// The delivery policy.
	// Default: - if the initial delivery of the message fails, three retries with a delay between failed attempts set at 20 seconds.
	//
	DeliveryPolicy *DeliveryPolicy `field:"optional" json:"deliveryPolicy" yaml:"deliveryPolicy"`
	// The filter policy.
	// Default: - all messages are delivered.
	//
	FilterPolicy *map[string]SubscriptionFilter `field:"optional" json:"filterPolicy" yaml:"filterPolicy"`
	// The filter policy that is applied on the message body.
	//
	// To apply a filter policy to the message attributes, use `filterPolicy`. A maximum of one of `filterPolicyWithMessageBody` and `filterPolicy` may be used.
	// Default: - all messages are delivered.
	//
	FilterPolicyWithMessageBody *map[string]FilterOrPolicy `field:"optional" json:"filterPolicyWithMessageBody" yaml:"filterPolicyWithMessageBody"`
	// true if raw message delivery is enabled for the subscription.
	//
	// Raw messages are free of JSON formatting and can be
	// sent to HTTP/S and Amazon SQS endpoints. For more information, see GetSubscriptionAttributes in the Amazon Simple
	// Notification Service API Reference.
	// Default: false.
	//
	RawMessageDelivery *bool `field:"optional" json:"rawMessageDelivery" yaml:"rawMessageDelivery"`
	// The region where the topic resides, in the case of cross-region subscriptions.
	// Default: - the region where the CloudFormation stack is being deployed.
	//
	Region *string `field:"optional" json:"region" yaml:"region"`
	// Arn of role allowing access to firehose delivery stream.
	//
	// Required for a firehose subscription protocol.
	// Default: - No subscription role is provided.
	//
	SubscriptionRoleArn *string `field:"optional" json:"subscriptionRoleArn" yaml:"subscriptionRoleArn"`
}

