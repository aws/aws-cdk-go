package awseventsv2


// The destination for events that could not be forwarded by the managed forwarding target or, for partner event sources, the managed partner event bus.
//
// Arn is optional. An empty object removes a configured destination on update.
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-eventsv2-eventsource-onfailureconfiguration.html
//
type CfnEventSource_OnFailureConfigurationProperty struct {
	// The ARN of the Amazon SQS standard queue that receives events that could not be forwarded.
	//
	// FIFO queues are not supported.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-eventsv2-eventsource-onfailureconfiguration.html#cfn-eventsv2-eventsource-onfailureconfiguration-arn
	//
	Arn *string `field:"optional" json:"arn" yaml:"arn"`
}

