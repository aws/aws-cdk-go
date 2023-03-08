package awssns


// Properties for defining a `CfnTopicPolicy`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var policyDocument interface{}
//
//   cfnTopicPolicyProps := &cfnTopicPolicyProps{
//   	policyDocument: policyDocument,
//   	topics: []*string{
//   		jsii.String("topics"),
//   	},
//   }
//
type CfnTopicPolicyProps struct {
	// A policy document that contains permissions to add to the specified SNS topics.
	PolicyDocument interface{} `field:"required" json:"policyDocument" yaml:"policyDocument"`
	// The Amazon Resource Names (ARN) of the topics to which you want to add the policy.
	//
	// You can use the `[Ref](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/intrinsic-function-reference-ref.html)` function to specify an `[AWS::SNS::Topic](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sns-topic.html)` resource.
	Topics *[]*string `field:"required" json:"topics" yaml:"topics"`
}

