package awssns

import (
	"github.com/aws/aws-cdk-go/awscdk/awskms"
)

// Properties for a new SNS topic.
//
// Example:
//   topic := sns.NewTopic(this, jsii.String("Topic"), &topicProps{
//   	displayName: jsii.String("Customer subscription topic"),
//   })
//
// Experimental.
type TopicProps struct {
	// Enables content-based deduplication for FIFO topics.
	// Experimental.
	ContentBasedDeduplication *bool `field:"optional" json:"contentBasedDeduplication" yaml:"contentBasedDeduplication"`
	// A developer-defined string that can be used to identify this SNS topic.
	// Experimental.
	DisplayName *string `field:"optional" json:"displayName" yaml:"displayName"`
	// Set to true to create a FIFO topic.
	// Experimental.
	Fifo *bool `field:"optional" json:"fifo" yaml:"fifo"`
	// A KMS Key, either managed by this CDK app, or imported.
	// Experimental.
	MasterKey awskms.IKey `field:"optional" json:"masterKey" yaml:"masterKey"`
	// A name for the topic.
	//
	// If you don't specify a name, AWS CloudFormation generates a unique
	// physical ID and uses that ID for the topic name. For more information,
	// see Name Type.
	// Experimental.
	TopicName *string `field:"optional" json:"topicName" yaml:"topicName"`
}

