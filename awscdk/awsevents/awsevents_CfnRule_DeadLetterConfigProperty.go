package awsevents


// A `DeadLetterConfig` object that contains information about a dead-letter queue configuration.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   deadLetterConfigProperty := &deadLetterConfigProperty{
//   	arn: jsii.String("arn"),
//   }
//
type CfnRule_DeadLetterConfigProperty struct {
	// The ARN of the SQS queue specified as the target for the dead-letter queue.
	Arn *string `field:"optional" json:"arn" yaml:"arn"`
}

