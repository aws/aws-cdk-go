package awssqs


// Properties for defining a `CfnQueueInlinePolicy`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var policyDocument interface{}
//
//   cfnQueueInlinePolicyProps := &CfnQueueInlinePolicyProps{
//   	PolicyDocument: policyDocument,
//   	Queue: jsii.String("queue"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sqs-queueinlinepolicy.html
//
type CfnQueueInlinePolicyProps struct {
	// A policy document that contains permissions to add to the specified SQS queue.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sqs-queueinlinepolicy.html#cfn-sqs-queueinlinepolicy-policydocument
	//
	PolicyDocument interface{} `field:"required" json:"policyDocument" yaml:"policyDocument"`
	// The URL of the SQS queue.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sqs-queueinlinepolicy.html#cfn-sqs-queueinlinepolicy-queue
	//
	Queue *string `field:"required" json:"queue" yaml:"queue"`
}

