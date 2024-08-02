package awssns

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awskms"
)

// Properties for a new SNS topic.
//
// Example:
//   topic := sns.NewTopic(this, jsii.String("MyTopic"), &TopicProps{
//   	TracingConfig: sns.TracingConfig_ACTIVE,
//   })
//
type TopicProps struct {
	// Enables content-based deduplication for FIFO topics.
	// Default: None.
	//
	ContentBasedDeduplication *bool `field:"optional" json:"contentBasedDeduplication" yaml:"contentBasedDeduplication"`
	// A developer-defined string that can be used to identify this SNS topic.
	//
	// The display name must be maximum 100 characters long, including hyphens (-),
	// underscores (_), spaces, and tabs.
	// Default: None.
	//
	DisplayName *string `field:"optional" json:"displayName" yaml:"displayName"`
	// Adds a statement to enforce encryption of data in transit when publishing to the topic.
	// See: https://docs.aws.amazon.com/sns/latest/dg/sns-security-best-practices.html#enforce-encryption-data-in-transit.
	//
	// Default: false.
	//
	EnforceSSL *bool `field:"optional" json:"enforceSSL" yaml:"enforceSSL"`
	// Set to true to create a FIFO topic.
	// Default: None.
	//
	Fifo *bool `field:"optional" json:"fifo" yaml:"fifo"`
	// The list of delivery status logging configurations for the topic.
	// See: https://docs.aws.amazon.com/sns/latest/dg/sns-topic-attributes.html.
	//
	// Default: None.
	//
	LoggingConfigs *[]*LoggingConfig `field:"optional" json:"loggingConfigs" yaml:"loggingConfigs"`
	// A KMS Key, either managed by this CDK app, or imported.
	// Default: None.
	//
	MasterKey awskms.IKey `field:"optional" json:"masterKey" yaml:"masterKey"`
	// The number of days Amazon SNS retains messages.
	//
	// It can only be set for FIFO topics.
	// See: https://docs.aws.amazon.com/sns/latest/dg/fifo-message-archiving-replay.html
	//
	// Default: - do not archive messages.
	//
	MessageRetentionPeriodInDays *float64 `field:"optional" json:"messageRetentionPeriodInDays" yaml:"messageRetentionPeriodInDays"`
	// The signature version corresponds to the hashing algorithm used while creating the signature of the notifications, subscription confirmations, or unsubscribe confirmation messages sent by Amazon SNS.
	// See: https://docs.aws.amazon.com/sns/latest/dg/sns-verify-signature-of-message.html.
	//
	// Default: 1.
	//
	SignatureVersion *string `field:"optional" json:"signatureVersion" yaml:"signatureVersion"`
	// A name for the topic.
	//
	// If you don't specify a name, AWS CloudFormation generates a unique
	// physical ID and uses that ID for the topic name. For more information,
	// see Name Type.
	// Default: Generated name.
	//
	TopicName *string `field:"optional" json:"topicName" yaml:"topicName"`
	// Tracing mode of an Amazon SNS topic.
	// See: https://docs.aws.amazon.com/sns/latest/dg/sns-active-tracing.html
	//
	// Default: TracingConfig.PASS_THROUGH
	//
	TracingConfig TracingConfig `field:"optional" json:"tracingConfig" yaml:"tracingConfig"`
}

