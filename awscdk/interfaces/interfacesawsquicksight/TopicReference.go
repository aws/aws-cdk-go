package interfacesawsquicksight


// A reference to a Topic resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   topicReference := &TopicReference{
//   	AwsAccountId: jsii.String("awsAccountId"),
//   	TopicArn: jsii.String("topicArn"),
//   	TopicId: jsii.String("topicId"),
//   }
//
type TopicReference struct {
	// The AwsAccountId of the Topic resource.
	AwsAccountId *string `field:"required" json:"awsAccountId" yaml:"awsAccountId"`
	// The ARN of the Topic resource.
	TopicArn *string `field:"required" json:"topicArn" yaml:"topicArn"`
	// The TopicId of the Topic resource.
	TopicId *string `field:"required" json:"topicId" yaml:"topicId"`
}

