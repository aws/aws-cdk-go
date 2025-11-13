package interfacesawsiot


// A reference to a TopicRule resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   topicRuleReference := &TopicRuleReference{
//   	RuleName: jsii.String("ruleName"),
//   	TopicRuleArn: jsii.String("topicRuleArn"),
//   }
//
type TopicRuleReference struct {
	// The RuleName of the TopicRule resource.
	RuleName *string `field:"required" json:"ruleName" yaml:"ruleName"`
	// The ARN of the TopicRule resource.
	TopicRuleArn *string `field:"required" json:"topicRuleArn" yaml:"topicRuleArn"`
}

