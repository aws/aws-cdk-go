package awssqs


// Properties for defining a `CfnQueuePolicy`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var policyDocument interface{}
//
//   cfnQueuePolicyProps := &cfnQueuePolicyProps{
//   	policyDocument: policyDocument,
//   	queues: []*string{
//   		jsii.String("queues"),
//   	},
//   }
//
type CfnQueuePolicyProps struct {
	// A policy document that contains the permissions for the specified Amazon SQS queues.
	//
	// For more information about Amazon SQS policies, see [Using custom policies with the Amazon SQS access policy language](https://docs.aws.amazon.com/AWSSimpleQueueService/latest/SQSDeveloperGuide/sqs-creating-custom-policies.html) in the *Amazon SQS Developer Guide* .
	PolicyDocument interface{} `field:"required" json:"policyDocument" yaml:"policyDocument"`
	// The URLs of the queues to which you want to add the policy.
	//
	// You can use the `[Ref](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/intrinsic-function-reference-ref.html)` function to specify an `[AWS::SQS::Queue](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sqs-queues.html)` resource.
	Queues *[]*string `field:"required" json:"queues" yaml:"queues"`
}

