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
//   eventDimensionsProperty := &EventDimensionsProperty{
//   	Attributes: attributes,
//   	EventType: &SetDimensionProperty{
//   		DimensionType: jsii.String("dimensionType"),
//   		Values: []*string{
//   			jsii.String("values"),
//   		},
//   	},
//   	Metrics: metrics,
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pinpoint-campaign-eventdimensions.html
//
type CfnCampaign_EventDimensionsProperty struct {
	// One or more custom attributes that your application reports to Amazon Pinpoint.
	//
	// You can use these attributes as selection criteria when you create an event filter.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pinpoint-campaign-eventdimensions.html#cfn-pinpoint-campaign-eventdimensions-attributes
	//
	Attributes interface{} `field:"optional" json:"attributes" yaml:"attributes"`
	// The name of the event that causes the campaign to be sent or the journey activity to be performed.
	//
	// This can be a standard event that Amazon Pinpoint generates, such as `_email.delivered` or `_custom.delivered` . For campaigns, this can also be a custom event that's specific to your application. For information about standard events, see [Streaming Amazon Pinpoint Events](https://docs.aws.amazon.com/pinpoint/latest/developerguide/event-streams.html) in the *Amazon Pinpoint Developer Guide* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pinpoint-campaign-eventdimensions.html#cfn-pinpoint-campaign-eventdimensions-eventtype
	//
	EventType interface{} `field:"optional" json:"eventType" yaml:"eventType"`
	// One or more custom metrics that your application reports to Amazon Pinpoint .
	//
	// You can use these metrics as selection criteria when you create an event filter.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pinpoint-campaign-eventdimensions.html#cfn-pinpoint-campaign-eventdimensions-metrics
	//
	Metrics interface{} `field:"optional" json:"metrics" yaml:"metrics"`
}

