package awssam


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   topicSAMPTProperty := &topicSAMPTProperty{
//   	topicName: jsii.String("topicName"),
//   }
//
type CfnFunction_TopicSAMPTProperty struct {
	// `CfnFunction.TopicSAMPTProperty.TopicName`.
	TopicName *string `field:"required" json:"topicName" yaml:"topicName"`
}

