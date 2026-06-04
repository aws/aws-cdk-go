package awssns


// The type of subscription, controlling the type of the endpoint parameter.
//
// Example:
//   // producerStack defines an SNS topic
//   var topic Topic
//
//
//   // consumerStack subscribes to it with a weak reference,
//   // so the producer can be torn down without blocking on this consumer
//   consumerStack := awscdk.Newstack(app, jsii.String("Consumer"), &StackProps{
//   	Env: &Environment{
//   		Account: jsii.String("123456789012"),
//   		Region: jsii.String("us-east-1"),
//   	},
//   })
//   sns.NewSubscription(consumerStack, jsii.String("Subscription"), &SubscriptionProps{
//   	Topic: sns.Topic_FromTopicArn(consumerStack, jsii.String("Topic"), awscdk.*stack_ConsumeReference(topic.topicArn)),
//   	Endpoint: jsii.String("https://example.com/webhook"),
//   	Protocol: sns.SubscriptionProtocol_HTTPS,
//   })
//
type SubscriptionProtocol string

const (
	// JSON-encoded message is POSTED to an HTTP url.
	SubscriptionProtocol_HTTP SubscriptionProtocol = "HTTP"
	// JSON-encoded message is POSTed to an HTTPS url.
	SubscriptionProtocol_HTTPS SubscriptionProtocol = "HTTPS"
	// Notifications are sent via email.
	SubscriptionProtocol_EMAIL SubscriptionProtocol = "EMAIL"
	// Notifications are JSON-encoded and sent via mail.
	SubscriptionProtocol_EMAIL_JSON SubscriptionProtocol = "EMAIL_JSON"
	// Notification is delivered by SMS.
	SubscriptionProtocol_SMS SubscriptionProtocol = "SMS"
	// Notifications are enqueued into an SQS queue.
	SubscriptionProtocol_SQS SubscriptionProtocol = "SQS"
	// JSON-encoded notifications are sent to a mobile app endpoint.
	SubscriptionProtocol_APPLICATION SubscriptionProtocol = "APPLICATION"
	// Notifications trigger a Lambda function.
	SubscriptionProtocol_LAMBDA SubscriptionProtocol = "LAMBDA"
	// Notifications put records into a firehose delivery stream.
	SubscriptionProtocol_FIREHOSE SubscriptionProtocol = "FIREHOSE"
)

