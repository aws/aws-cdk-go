package previewawssnsmixins


// Properties for CfnTopicInlinePolicyPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   var policyDocument interface{}
//
//   cfnTopicInlinePolicyMixinProps := &CfnTopicInlinePolicyMixinProps{
//   	PolicyDocument: policyDocument,
//   	TopicArn: jsii.String("topicArn"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sns-topicinlinepolicy.html
//
type CfnTopicInlinePolicyMixinProps struct {
	// A policy document that contains permissions to add to the specified Amazon  topic.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sns-topicinlinepolicy.html#cfn-sns-topicinlinepolicy-policydocument
	//
	PolicyDocument interface{} `field:"optional" json:"policyDocument" yaml:"policyDocument"`
	// The Amazon Resource Name (ARN) of the topic to which you want to add the policy.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sns-topicinlinepolicy.html#cfn-sns-topicinlinepolicy-topicarn
	//
	TopicArn *string `field:"optional" json:"topicArn" yaml:"topicArn"`
}

