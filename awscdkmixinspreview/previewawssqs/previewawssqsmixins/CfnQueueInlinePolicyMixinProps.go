package previewawssqsmixins


// Properties for CfnQueueInlinePolicyPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   var policyDocument interface{}
//
//   cfnQueueInlinePolicyMixinProps := &CfnQueueInlinePolicyMixinProps{
//   	PolicyDocument: policyDocument,
//   	Queue: jsii.String("queue"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sqs-queueinlinepolicy.html
//
type CfnQueueInlinePolicyMixinProps struct {
	// A policy document that contains the permissions for the specified Amazon SQS queues.
	//
	// For more information about Amazon SQS policies, see [Using custom policies with the Amazon SQS access policy language](https://docs.aws.amazon.com/AWSSimpleQueueService/latest/SQSDeveloperGuide/sqs-creating-custom-policies.html) in the *Amazon SQS Developer Guide* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sqs-queueinlinepolicy.html#cfn-sqs-queueinlinepolicy-policydocument
	//
	PolicyDocument interface{} `field:"optional" json:"policyDocument" yaml:"policyDocument"`
	// The URLs of the queues to which you want to add the policy.
	//
	// You can use the `[Ref](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/intrinsic-function-reference-ref.html)` function to specify an `[AWS::SQS::Queue](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sqs-queues.html)` resource.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sqs-queueinlinepolicy.html#cfn-sqs-queueinlinepolicy-queue
	//
	Queue *string `field:"optional" json:"queue" yaml:"queue"`
}

