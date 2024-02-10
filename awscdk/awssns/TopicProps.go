package awssns

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awskms"
)

// Properties for a new SNS topic.
//
// Example:
//   var role role
//
//   topic := sns.NewTopic(this, jsii.String("MyTopic"), &TopicProps{
//   	Fifo: jsii.Boolean(true),
//   	MessageRetentionPeriodInDays: jsii.Number(7),
//   })
//
type TopicProps struct {
	// Enables content-based deduplication for FIFO topics.
	// Default: None.
	//
	ContentBasedDeduplication *bool `field:"optional" json:"contentBasedDeduplication" yaml:"contentBasedDeduplication"`
	// A developer-defined string that can be used to identify this SNS topic.
	// Default: None.
	//
	DisplayName *string `field:"optional" json:"displayName" yaml:"displayName"`
	// Set to true to create a FIFO topic.
	// Default: None.
	//
	Fifo *bool `field:"optional" json:"fifo" yaml:"fifo"`
	// The list of delivery status logging configurations for the topic.
	//
	// For more information, see https://docs.aws.amazon.com/sns/latest/dg/sns-topic-attributes.html.
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
	// A name for the topic.
	//
	// If you don't specify a name, AWS CloudFormation generates a unique
	// physical ID and uses that ID for the topic name. For more information,
	// see Name Type.
	// Default: Generated name.
	//
	TopicName *string `field:"optional" json:"topicName" yaml:"topicName"`
}

