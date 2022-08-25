package awspinpoint


// Specifies the dimensions for an event filter that determines when a campaign is sent or a journey activity is performed.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var attributes interface{}
//   var metrics interface{}
//
//   eventDimensionsProperty := &eventDimensionsProperty{
//   	attributes: attributes,
//   	eventType: &setDimensionProperty{
//   		dimensionType: jsii.String("dimensionType"),
//   		values: []*string{
//   			jsii.String("values"),
//   		},
//   	},
//   	metrics: metrics,
//   }
//
type CfnCampaign_EventDimensionsProperty struct {
	// One or more custom attributes that your application reports to Amazon Pinpoint.
	//
	// You can use these attributes as selection criteria when you create an event filter.
	Attributes interface{} `field:"optional" json:"attributes" yaml:"attributes"`
	// The name of the event that causes the campaign to be sent or the journey activity to be performed.
	//
	// This can be a standard event that Amazon Pinpoint generates, such as `_email.delivered` . For campaigns, this can also be a custom event that's specific to your application. For information about standard events, see [Streaming Amazon Pinpoint Events](https://docs.aws.amazon.com/pinpoint/latest/developerguide/event-streams.html) in the *Amazon Pinpoint Developer Guide* .
	EventType interface{} `field:"optional" json:"eventType" yaml:"eventType"`
	// One or more custom metrics that your application reports to Amazon Pinpoint .
	//
	// You can use these metrics as selection criteria when you create an event filter.
	Metrics interface{} `field:"optional" json:"metrics" yaml:"metrics"`
}

