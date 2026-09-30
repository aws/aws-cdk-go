package awsquicksight


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   topicReferenceProperty := &TopicReferenceProperty{
//   	TopicArn: jsii.String("topicArn"),
//   	TopicPlaceholder: jsii.String("topicPlaceholder"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-topicreference.html
//
type CfnTemplatePropsMixin_TopicReferenceProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-topicreference.html#cfn-quicksight-template-topicreference-topicarn
	//
	TopicArn *string `field:"optional" json:"topicArn" yaml:"topicArn"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-template-topicreference.html#cfn-quicksight-template-topicreference-topicplaceholder
	//
	TopicPlaceholder *string `field:"optional" json:"topicPlaceholder" yaml:"topicPlaceholder"`
}

