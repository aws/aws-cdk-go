package awssns


// `Subscription` is an embedded property that describes the subscription endpoints of an Amazon SNS topic.
//
// > For full control over subscription behavior (for example, delivery policy, filtering, raw message delivery, and cross-region subscriptions), use the [AWS::SNS::Subscription](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sns-subscription.html) resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   subscriptionProperty := &subscriptionProperty{
//   	endpoint: jsii.String("endpoint"),
//   	protocol: jsii.String("protocol"),
//   }
//
type CfnTopic_SubscriptionProperty struct {
	// The endpoint that receives notifications from the Amazon SNS topic.
	//
	// The endpoint value depends on the protocol that you specify. For more information, see the `Endpoint` parameter of the `[Subscribe](https://docs.aws.amazon.com/sns/latest/api/API_Subscribe.html)` action in the *Amazon SNS API Reference* .
	Endpoint *string `field:"required" json:"endpoint" yaml:"endpoint"`
	// The subscription's protocol.
	//
	// For more information, see the `Protocol` parameter of the `[Subscribe](https://docs.aws.amazon.com/sns/latest/api/API_Subscribe.html)` action in the *Amazon SNS API Reference* .
	Protocol *string `field:"required" json:"protocol" yaml:"protocol"`
}

