package awsevents


// Properties for defining a `CfnEventBus`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnEventBusProps := &CfnEventBusProps{
//   	Name: jsii.String("name"),
//
//   	// the properties below are optional
//   	EventSourceName: jsii.String("eventSourceName"),
//   	Tags: []tagEntryProperty{
//   		&tagEntryProperty{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-events-eventbus.html
//
type CfnEventBusProps struct {
	// The name of the new event bus.
	//
	// Custom event bus names can't contain the `/` character, but you can use the `/` character in partner event bus names. In addition, for partner event buses, the name must exactly match the name of the partner event source that this event bus is matched to.
	//
	// You can't use the name `default` for a custom event bus, as this name is already used for your account's default event bus.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-events-eventbus.html#cfn-events-eventbus-name
	//
	Name *string `field:"required" json:"name" yaml:"name"`
	// If you are creating a partner event bus, this specifies the partner event source that the new event bus will be matched with.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-events-eventbus.html#cfn-events-eventbus-eventsourcename
	//
	EventSourceName *string `field:"optional" json:"eventSourceName" yaml:"eventSourceName"`
	// Tags to associate with the event bus.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-events-eventbus.html#cfn-events-eventbus-tags
	//
	Tags *[]*CfnEventBus_TagEntryProperty `field:"optional" json:"tags" yaml:"tags"`
}

