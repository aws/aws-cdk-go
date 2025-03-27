package awsevents


// Configuration details of the Amazon SQS queue for EventBridge to use as a dead-letter queue (DLQ).
//
// For more information, see [Using dead-letter queues to process undelivered events](https://docs.aws.amazon.com/eventbridge/latest/userguide/eb-rule-event-delivery.html#eb-rule-dlq) in the *EventBridge User Guide* .
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-events-eventbus-deadletterconfig.html
//
type CfnEventBus_DeadLetterConfigProperty struct {
	// The ARN of the SQS queue specified as the target for the dead-letter queue.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-events-eventbus-deadletterconfig.html#cfn-events-eventbus-deadletterconfig-arn
	//
	Arn *string `field:"optional" json:"arn" yaml:"arn"`
}

