package awseventsv2


// The destination for events that could not be delivered.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   onFailureConfigurationProperty := &OnFailureConfigurationProperty{
//   	Arn: jsii.String("arn"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-eventsv2-subscriber-onfailureconfiguration.html
//
type CfnSubscriber_OnFailureConfigurationProperty struct {
	// The ARN of the destination that receives events that could not be delivered.
	//
	// An Amazon SQS queue is the supported destination.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-eventsv2-subscriber-onfailureconfiguration.html#cfn-eventsv2-subscriber-onfailureconfiguration-arn
	//
	Arn *string `field:"optional" json:"arn" yaml:"arn"`
}

