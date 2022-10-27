package awssns


// The type of subscription, controlling the type of the endpoint parameter.
//
// Example:
//   // Example automatically generated from non-compiling source. May contain errors.
//   import "github.com/aws-samples/dummy/awscdklibawskinesisfirehose"
//   var stream DeliveryStream
//
//
//   topic := sns.NewTopic(this, jsii.String("Topic"))
//
//   sns.NewSubscription(this, jsii.String("Subscription"), &subscriptionProps{
//   	topic: topic,
//   	endpoint: stream.deliveryStreamArn,
//   	protocol: sns.subscriptionProtocol_FIREHOSE,
//   	subscriptionRoleArn: jsii.String("SAMPLE_ARN"),
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

