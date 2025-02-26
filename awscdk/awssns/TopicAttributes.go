package awssns


// Represents an SNS topic defined outside of this stack.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   topicAttributes := &TopicAttributes{
//   	TopicArn: jsii.String("topicArn"),
//
//   	// the properties below are optional
//   	ContentBasedDeduplication: jsii.Boolean(false),
//   }
//
type TopicAttributes struct {
	// The ARN of the SNS topic.
	TopicArn *string `field:"required" json:"topicArn" yaml:"topicArn"`
	// Whether content-based deduplication is enabled.
	//
	// Only applicable for FIFO topics.
	// Default: false.
	//
	ContentBasedDeduplication *bool `field:"optional" json:"contentBasedDeduplication" yaml:"contentBasedDeduplication"`
}

