package awssns

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awssqs"
	"github.com/aws/constructs-go/constructs/v10"
)

// Subscription configuration.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import constructs "github.com/aws/constructs-go/constructs"
//
//   var construct construct
//   var dependable iDependable
//   var filterOrPolicy filterOrPolicy
//   var queue queue
//   var subscriptionFilter subscriptionFilter
//
//   topicSubscriptionConfig := &TopicSubscriptionConfig{
//   	Endpoint: jsii.String("endpoint"),
//   	Protocol: awscdk.Aws_sns.SubscriptionProtocol_HTTP,
//   	SubscriberId: jsii.String("subscriberId"),
//
//   	// the properties below are optional
//   	DeadLetterQueue: queue,
//   	FilterPolicy: map[string]*subscriptionFilter{
//   		"filterPolicyKey": subscriptionFilter,
//   	},
//   	FilterPolicyWithMessageBody: map[string]*filterOrPolicy{
//   		"filterPolicyWithMessageBodyKey": filterOrPolicy,
//   	},
//   	RawMessageDelivery: jsii.Boolean(false),
//   	Region: jsii.String("region"),
//   	SubscriberScope: construct,
//   	SubscriptionDependency: dependable,
//   	SubscriptionRoleArn: jsii.String("subscriptionRoleArn"),
//   }
//
type TopicSubscriptionConfig struct {
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
	// The id of the SNS subscription resource created under `scope`.
	//
	// In most
	// cases, it is recommended to use the `uniqueId` of the topic you are
	// subscribing to.
	SubscriberId *string `field:"required" json:"subscriberId" yaml:"subscriberId"`
	// The scope in which to create the SNS subscription resource.
	//
	// Normally you'd
	// want the subscription to be created on the consuming stack because the
	// topic is usually referenced by the consumer's resource policy (e.g. SQS
	// queue policy). Otherwise, it will cause a cyclic reference.
	//
	// If this is undefined, the subscription will be created on the topic's stack.
	// Default: - use the topic as the scope of the subscription, in which case `subscriberId` must be defined.
	//
	SubscriberScope constructs.Construct `field:"optional" json:"subscriberScope" yaml:"subscriberScope"`
	// The resources that need to be created before the subscription can be safely created.
	//
	// For example for SQS subscription, the subscription needs to have a dependency on the SQS queue policy
	// in order for the subscription to successfully deliver messages.
	// Default: - empty list.
	//
	SubscriptionDependency constructs.IDependable `field:"optional" json:"subscriptionDependency" yaml:"subscriptionDependency"`
}

