package awspipes


// A `DeadLetterConfig` object that contains information about a dead-letter queue configuration.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   deadLetterConfigProperty := &DeadLetterConfigProperty{
//   	Arn: jsii.String("arn"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pipes-pipe-deadletterconfig.html
//
type CfnPipe_DeadLetterConfigProperty struct {
	// The ARN of the specified target for the dead-letter queue.
	//
	// For Amazon Kinesis stream and Amazon DynamoDB stream sources, specify either an Amazon SNS topic or Amazon SQS queue ARN.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pipes-pipe-deadletterconfig.html#cfn-pipes-pipe-deadletterconfig-arn
	//
	Arn *string `field:"optional" json:"arn" yaml:"arn"`
}

