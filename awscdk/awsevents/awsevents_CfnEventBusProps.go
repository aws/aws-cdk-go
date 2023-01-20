package awsevents


// Properties for defining a `CfnEventBus`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnEventBusProps := &cfnEventBusProps{
//   	name: jsii.String("name"),
//
//   	// the properties below are optional
//   	eventSourceName: jsii.String("eventSourceName"),
//   	tags: []tagEntryProperty{
//   		&tagEntryProperty{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   }
//
type CfnEventBusProps struct {
	// The name of the new event bus.
	//
	// Event bus names cannot contain the / character. You can't use the name `default` for a custom event bus, as this name is already used for your account's default event bus.
	//
	// If this is a partner event bus, the name must exactly match the name of the partner event source that this event bus is matched to.
	Name *string `field:"required" json:"name" yaml:"name"`
	// If you are creating a partner event bus, this specifies the partner event source that the new event bus will be matched with.
	EventSourceName *string `field:"optional" json:"eventSourceName" yaml:"eventSourceName"`
	// Tags to associate with the event bus.
	Tags *[]*CfnEventBus_TagEntryProperty `field:"optional" json:"tags" yaml:"tags"`
}

