package awseventsv2


// Customer-controllable system metadata attached to each forwarded event.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   eventBusV2SystemMetadataProperty := &EventBusV2SystemMetadataProperty{
//   	DeduplicationId: jsii.String("deduplicationId"),
//   	EventGroupId: jsii.String("eventGroupId"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-eventsv2-subscriber-eventbusv2systemmetadata.html
//
type CfnSubscriberPropsMixin_EventBusV2SystemMetadataProperty struct {
	// The deduplication ID for FIFO deduplication on the downstream event bus.
	//
	// Accepts a literal value or a JSONata expression.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-eventsv2-subscriber-eventbusv2systemmetadata.html#cfn-eventsv2-subscriber-eventbusv2systemmetadata-deduplicationid
	//
	DeduplicationId *string `field:"optional" json:"deduplicationId" yaml:"deduplicationId"`
	// The event group ID for FIFO ordering on the downstream event bus.
	//
	// Accepts a literal value or a JSONata expression.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-eventsv2-subscriber-eventbusv2systemmetadata.html#cfn-eventsv2-subscriber-eventbusv2systemmetadata-eventgroupid
	//
	EventGroupId *string `field:"optional" json:"eventGroupId" yaml:"eventGroupId"`
}

