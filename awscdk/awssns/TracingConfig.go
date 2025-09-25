package awssns


// The tracing mode of an Amazon SNS topic.
//
// Example:
//   topic := sns.NewTopic(this, jsii.String("MyTopic"), &TopicProps{
//   	TracingConfig: sns.TracingConfig_ACTIVE,
//   })
//
type TracingConfig string

const (
	// The mode that topic passes trace headers received from the Amazon SNS publisher to its subscription.
	TracingConfig_PASS_THROUGH TracingConfig = "PASS_THROUGH"
	// The mode that Amazon SNS vend X-Ray segment data to topic owner account if the sampled flag in the tracing header is true.
	TracingConfig_ACTIVE TracingConfig = "ACTIVE"
)

