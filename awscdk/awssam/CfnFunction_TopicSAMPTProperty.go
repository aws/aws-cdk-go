package awssam


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   topicSAMPTProperty := &TopicSAMPTProperty{
//   	TopicName: jsii.String("topicName"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-serverless-function-topicsampt.html
//
type CfnFunction_TopicSAMPTProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-serverless-function-topicsampt.html#cfn-serverless-function-topicsampt-topicname
	//
	TopicName *string `field:"required" json:"topicName" yaml:"topicName"`
}

