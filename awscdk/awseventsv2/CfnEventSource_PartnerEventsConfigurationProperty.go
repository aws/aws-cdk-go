package awseventsv2


// Forward a partner event source's events through a managed partner event bus.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   partnerEventsConfigurationProperty := &PartnerEventsConfigurationProperty{
//   	PartnerEventSourceArn: jsii.String("partnerEventSourceArn"),
//
//   	// the properties below are optional
//   	OnFailureConfiguration: &OnFailureConfigurationProperty{
//   		Arn: jsii.String("arn"),
//   	},
//   	PartnerBusKmsKeyIdentifier: jsii.String("partnerBusKmsKeyIdentifier"),
//   	Pattern: jsii.String("pattern"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-eventsv2-eventsource-partnereventsconfiguration.html
//
type CfnEventSource_PartnerEventsConfigurationProperty struct {
	// The ARN of the partner event source to forward.
	//
	// The partner owns the event source, so the ARN's account segment is empty. Changing this property replaces the event source. Because Name and EventBusArn together identify an event source, and the replacement is created before the old resource is deleted, change Name in the same update.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-eventsv2-eventsource-partnereventsconfiguration.html#cfn-eventsv2-eventsource-partnereventsconfiguration-partnereventsourcearn
	//
	PartnerEventSourceArn *string `field:"required" json:"partnerEventSourceArn" yaml:"partnerEventSourceArn"`
	// The destination for events that could not be forwarded by the managed forwarding target or, for partner event sources, the managed partner event bus.
	//
	// Arn is optional. An empty object removes a configured destination on update.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-eventsv2-eventsource-partnereventsconfiguration.html#cfn-eventsv2-eventsource-partnereventsconfiguration-onfailureconfiguration
	//
	OnFailureConfiguration interface{} `field:"optional" json:"onFailureConfiguration" yaml:"onFailureConfiguration"`
	// The identifier of the AWS KMS customer managed key for EventBridge to use, if you choose to use a customer managed key to encrypt events on the managed partner event bus.
	//
	// The identifier can be the key Amazon Resource Name (ARN), KeyId, key alias, or key alias ARN. If you do not specify a customer managed key identifier, EventBridge uses an AWS owned key to encrypt events on the event bus.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-eventsv2-eventsource-partnereventsconfiguration.html#cfn-eventsv2-eventsource-partnereventsconfiguration-partnerbuskmskeyidentifier
	//
	PartnerBusKmsKeyIdentifier *string `field:"optional" json:"partnerBusKmsKeyIdentifier" yaml:"partnerBusKmsKeyIdentifier"`
	// A filter pattern, as a JSON string, that defines which events from the specified partner event source are forwarded to the event bus.
	//
	// If you do not specify a pattern, all events from the partner event source are forwarded.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-eventsv2-eventsource-partnereventsconfiguration.html#cfn-eventsv2-eventsource-partnereventsconfiguration-pattern
	//
	Pattern *string `field:"optional" json:"pattern" yaml:"pattern"`
}

