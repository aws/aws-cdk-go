package awssns


// Options for customising the delivery of SNS messages to HTTP/S endpoints.
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
type DeliveryPolicy struct {
	// The retry policy of the delivery of SNS messages to HTTP/S endpoints.
	// Default: - Amazon SNS attempts up to three retries with a delay between failed attempts set at 20 seconds.
	//
	HealthyRetryPolicy *HealthyRetryPolicy `field:"optional" json:"healthyRetryPolicy" yaml:"healthyRetryPolicy"`
	// The request of the content sent in AWS SNS HTTP/S requests.
	// Default: - The content type is set to 'text/plain; charset=UTF-8'.
	//
	RequestPolicy *RequestPolicy `field:"optional" json:"requestPolicy" yaml:"requestPolicy"`
	// The throttling policy of the delivery of SNS messages to HTTP/S endpoints.
	// Default: - No throttling.
	//
	ThrottlePolicy *ThrottlePolicy `field:"optional" json:"throttlePolicy" yaml:"throttlePolicy"`
}

