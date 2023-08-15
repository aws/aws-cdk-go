package awssns


// Properties for defining a `CfnTopicInlinePolicy`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var policyDocument interface{}
//
//   cfnTopicInlinePolicyProps := &CfnTopicInlinePolicyProps{
//   	PolicyDocument: policyDocument,
//   	TopicArn: jsii.String("topicArn"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sns-topicinlinepolicy.html
//
type CfnTopicInlinePolicyProps struct {
	// A policy document that contains permissions to add to the specified Amazon SNS topic.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sns-topicinlinepolicy.html#cfn-sns-topicinlinepolicy-policydocument
	//
	PolicyDocument interface{} `field:"required" json:"policyDocument" yaml:"policyDocument"`
	// The Amazon Resource Name (ARN) of the topic to which you want to add the policy.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sns-topicinlinepolicy.html#cfn-sns-topicinlinepolicy-topicarn
	//
	TopicArn *string `field:"required" json:"topicArn" yaml:"topicArn"`
}

