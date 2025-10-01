package awsiot


// A reference to a TopicRuleDestination resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   topicRuleDestinationReference := &TopicRuleDestinationReference{
//   	TopicRuleDestinationArn: jsii.String("topicRuleDestinationArn"),
//   }
//
type TopicRuleDestinationReference struct {
	// The Arn of the TopicRuleDestination resource.
	TopicRuleDestinationArn *string `field:"required" json:"topicRuleDestinationArn" yaml:"topicRuleDestinationArn"`
}

