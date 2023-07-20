package awsevents


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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-events-rule-deadletterconfig.html
//
type CfnRule_DeadLetterConfigProperty struct {
	// The ARN of the SQS queue specified as the target for the dead-letter queue.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-events-rule-deadletterconfig.html#cfn-events-rule-deadletterconfig-arn
	//
	Arn *string `field:"optional" json:"arn" yaml:"arn"`
}

