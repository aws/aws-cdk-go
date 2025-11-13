package interfacesawssns


// A reference to a TopicPolicy resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   topicPolicyReference := &TopicPolicyReference{
//   	TopicPolicyId: jsii.String("topicPolicyId"),
//   }
//
type TopicPolicyReference struct {
	// The Id of the TopicPolicy resource.
	TopicPolicyId *string `field:"required" json:"topicPolicyId" yaml:"topicPolicyId"`
}

