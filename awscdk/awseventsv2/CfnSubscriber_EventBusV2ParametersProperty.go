package awseventsv2


// Parameters for forwarding events to another EventBridge event bus, used when TargetArn is an event bus ARN of the form arn:{partition}:events:{region}:{account}:event-busv2/{name}/{id}.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   eventBusV2ParametersProperty := &EventBusV2ParametersProperty{
//   	DeduplicationConfiguration: &DeduplicationConfigurationProperty{
//   		DeduplicationType: jsii.String("deduplicationType"),
//   	},
//   	Metadata: map[string]*string{
//   		"metadataKey": jsii.String("metadata"),
//   	},
//   	SystemMetadata: &EventBusV2SystemMetadataProperty{
//   		DeduplicationId: jsii.String("deduplicationId"),
//   		EventGroupId: jsii.String("eventGroupId"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-eventsv2-subscriber-eventbusv2parameters.html
//
type CfnSubscriber_EventBusV2ParametersProperty struct {
	// Deduplication settings applied to the forwarded events on the downstream bus.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-eventsv2-subscriber-eventbusv2parameters.html#cfn-eventsv2-subscriber-eventbusv2parameters-deduplicationconfiguration
	//
	DeduplicationConfiguration interface{} `field:"optional" json:"deduplicationConfiguration" yaml:"deduplicationConfiguration"`
	// Metadata forwarded with each event, as key-value string pairs, with at most 100 entries.
	//
	// A key or value may also be a JSONata expression.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-eventsv2-subscriber-eventbusv2parameters.html#cfn-eventsv2-subscriber-eventbusv2parameters-metadata
	//
	Metadata interface{} `field:"optional" json:"metadata" yaml:"metadata"`
	// Customer-controllable system metadata attached to each forwarded event.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-eventsv2-subscriber-eventbusv2parameters.html#cfn-eventsv2-subscriber-eventbusv2parameters-systemmetadata
	//
	SystemMetadata interface{} `field:"optional" json:"systemMetadata" yaml:"systemMetadata"`
}

