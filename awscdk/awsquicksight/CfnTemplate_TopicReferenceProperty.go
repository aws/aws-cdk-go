package awsquicksight


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   topicReferenceProperty := &TopicReferenceProperty{
//   	TopicArn: jsii.String("topicArn"),
//   	TopicPlaceholder: jsii.String("topicPlaceholder"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-topicreference.html
//
type CfnTemplate_TopicReferenceProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-topicreference.html#cfn-quicksight-template-topicreference-topicarn
	//
	TopicArn *string `field:"required" json:"topicArn" yaml:"topicArn"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-topicreference.html#cfn-quicksight-template-topicreference-topicplaceholder
	//
	TopicPlaceholder *string `field:"required" json:"topicPlaceholder" yaml:"topicPlaceholder"`
}

