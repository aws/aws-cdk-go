package awssns


// The type of supported protocol for delivery status logging.
//
// Example:
//   var role role
//
//   topic := sns.NewTopic(this, jsii.String("MyTopic"), &TopicProps{
//   	LoggingConfigs: []loggingConfig{
//   		&loggingConfig{
//   			Protocol: sns.LoggingProtocol_SQS,
//   			FailureFeedbackRole: role,
//   			SuccessFeedbackRole: role,
//   			SuccessFeedbackSampleRate: jsii.Number(50),
//   		},
//   	},
//   })
//
type LoggingProtocol string

const (
	// HTTP.
	LoggingProtocol_HTTP LoggingProtocol = "HTTP"
	// Amazon Simple Queue Service.
	LoggingProtocol_SQS LoggingProtocol = "SQS"
	// AWS Lambda.
	LoggingProtocol_LAMBDA LoggingProtocol = "LAMBDA"
	// Amazon Kinesis Data Firehose.
	LoggingProtocol_FIREHOSE LoggingProtocol = "FIREHOSE"
	// Platform application endpoint.
	LoggingProtocol_APPLICATION LoggingProtocol = "APPLICATION"
)

