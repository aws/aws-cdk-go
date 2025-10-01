package awssns


// Options for customising aspects of the content sent in AWS SNS HTTP/S requests.
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
type RequestPolicy struct {
	// The content type of the notification being sent to HTTP/S endpoints.
	// Default: - text/plain; charset=UTF-8.
	//
	HeaderContentType *string `field:"optional" json:"headerContentType" yaml:"headerContentType"`
}

