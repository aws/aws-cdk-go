package interfacesawssns


// A reference to a TopicInlinePolicy resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   topicInlinePolicyReference := &TopicInlinePolicyReference{
//   	TopicArn: jsii.String("topicArn"),
//   }
//
type TopicInlinePolicyReference struct {
	// The TopicArn of the TopicInlinePolicy resource.
	TopicArn *string `field:"required" json:"topicArn" yaml:"topicArn"`
}

