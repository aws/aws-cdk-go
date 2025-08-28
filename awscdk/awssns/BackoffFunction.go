package awssns


// Algorithms which can be used by SNS to calculate the delays associated with all of the retry attempts between the first and last retries in the backoff phase.
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
type BackoffFunction string

const (
	// Arithmetic, see {@link https://docs.aws.amazon.com/images/sns/latest/dg/images/backoff-graph.png this image} for how this function compares to others.
	BackoffFunction_ARITHMETIC BackoffFunction = "ARITHMETIC"
	// Exponential, see {@link https://docs.aws.amazon.com/images/sns/latest/dg/images/backoff-graph.png this image} for how this function compares to others.
	BackoffFunction_EXPONENTIAL BackoffFunction = "EXPONENTIAL"
	// Geometric, see {@link https://docs.aws.amazon.com/images/sns/latest/dg/images/backoff-graph.png this image} for how this function compares to others.
	BackoffFunction_GEOMETRIC BackoffFunction = "GEOMETRIC"
	// Linear, see {@link https://docs.aws.amazon.com/images/sns/latest/dg/images/backoff-graph.png this image} for how this function compares to others.
	BackoffFunction_LINEAR BackoffFunction = "LINEAR"
)

