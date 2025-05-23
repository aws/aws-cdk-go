package awssns

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Options for customising the retry policy of the delivery of SNS messages to HTTP/S endpoints.
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
type HealthyRetryPolicy struct {
	// The model for backoff between retries.
	// Default: - linear.
	//
	BackoffFunction BackoffFunction `field:"optional" json:"backoffFunction" yaml:"backoffFunction"`
	// The maximum delay for a retry.
	//
	// Must be at least `minDelayTarget` less than 3,600 seconds, and correspond to a whole number of seconds,.
	// Default: - 20 seconds.
	//
	MaxDelayTarget awscdk.Duration `field:"optional" json:"maxDelayTarget" yaml:"maxDelayTarget"`
	// The minimum delay for a retry.
	//
	// Must be at least one second, not exceed `maxDelayTarget`, and correspond to a whole number of seconds.
	// Default: - 20 seconds.
	//
	MinDelayTarget awscdk.Duration `field:"optional" json:"minDelayTarget" yaml:"minDelayTarget"`
	// The number of retries in the post-backoff phase, with the maximum delay between them.
	//
	// Must be zero or greater.
	// Default: 0.
	//
	NumMaxDelayRetries *float64 `field:"optional" json:"numMaxDelayRetries" yaml:"numMaxDelayRetries"`
	// The number of retries in the pre-backoff phase, with the specified minimum delay between them.
	//
	// Must be zero or greater.
	// Default: 0.
	//
	NumMinDelayRetries *float64 `field:"optional" json:"numMinDelayRetries" yaml:"numMinDelayRetries"`
	// The number of retries to be done immediately, with no delay between them.
	//
	// Must be zero or greater.
	// Default: 0.
	//
	NumNoDelayRetries *float64 `field:"optional" json:"numNoDelayRetries" yaml:"numNoDelayRetries"`
	// The total number of retries, including immediate, pre-backoff, backoff, and post-backoff retries.
	//
	// Must be greater than or equal to zero and not exceed 100.
	// Default: 3.
	//
	NumRetries *float64 `field:"optional" json:"numRetries" yaml:"numRetries"`
}

