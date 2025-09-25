package awssns


// A reference to a Topic resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   topicReference := &TopicReference{
//   	TopicArn: jsii.String("topicArn"),
//   }
//
type TopicReference struct {
	// The TopicArn of the Topic resource.
	TopicArn *string `field:"required" json:"topicArn" yaml:"topicArn"`
}

