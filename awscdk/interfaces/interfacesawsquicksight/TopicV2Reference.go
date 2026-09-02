package interfacesawsquicksight


// A reference to a TopicV2 resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   topicV2Reference := &TopicV2Reference{
//   	AwsAccountId: jsii.String("awsAccountId"),
//   	TopicId: jsii.String("topicId"),
//   	TopicV2Arn: jsii.String("topicV2Arn"),
//   }
//
type TopicV2Reference struct {
	// The AwsAccountId of the TopicV2 resource.
	AwsAccountId *string `field:"required" json:"awsAccountId" yaml:"awsAccountId"`
	// The TopicId of the TopicV2 resource.
	TopicId *string `field:"required" json:"topicId" yaml:"topicId"`
	// The ARN of the TopicV2 resource.
	TopicV2Arn *string `field:"required" json:"topicV2Arn" yaml:"topicV2Arn"`
}

