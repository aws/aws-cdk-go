package awsevents


// Dead Letter Queue for the event bus.
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
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-events-eventbus-deadletterconfig.html#cfn-events-eventbus-deadletterconfig-arn
	//
	Arn *string `field:"optional" json:"arn" yaml:"arn"`
}

