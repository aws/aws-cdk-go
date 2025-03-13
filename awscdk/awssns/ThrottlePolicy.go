package awssns


// Options for customising AWS SNS HTTP/S delivery throttling.
//
// Example:
//   myTopic := sns.NewTopic(this, jsii.String("MyTopic"))
//
//   myTopic.AddSubscription(
//   subscriptions.NewUrlSubscription(jsii.String("https://foobar.com/"), &UrlSubscriptionProps{
//   	DeliveryPolicy: &DeliveryPolicy{
//   		HealthyRetryPolicy: &HealthyRetryPolicy{
//   			MinDelayTarget: awscdk.Duration_Seconds(jsii.Number(5)),
//   			MaxDelayTarget: awscdk.Duration_*Seconds(jsii.Number(10)),
//   			NumRetries: jsii.Number(6),
//   			BackoffFunction: sns.BackoffFunction_EXPONENTIAL,
//   		},
//   		ThrottlePolicy: &ThrottlePolicy{
//   			MaxReceivesPerSecond: jsii.Number(10),
//   		},
//   		RequestPolicy: &RequestPolicy{
//   			HeaderContentType: jsii.String("application/json"),
//   		},
//   	},
//   }))
//
type ThrottlePolicy struct {
	// The maximum number of deliveries per second, per subscription.
	// Default: - no throttling.
	//
	MaxReceivesPerSecond *float64 `field:"optional" json:"maxReceivesPerSecond" yaml:"maxReceivesPerSecond"`
}

