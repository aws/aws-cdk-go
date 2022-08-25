package awslambda


// The [dead-letter queue](https://docs.aws.amazon.com/lambda/latest/dg/invocation-async.html#dlq) for failed asynchronous invocations.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   deadLetterConfigProperty := &deadLetterConfigProperty{
//   	targetArn: jsii.String("targetArn"),
//   }
//
type CfnFunction_DeadLetterConfigProperty struct {
	// The Amazon Resource Name (ARN) of an Amazon SQS queue or Amazon SNS topic.
	TargetArn *string `field:"optional" json:"targetArn" yaml:"targetArn"`
}

