package awspinpoint


// Specifies the settings for events that cause a campaign to be sent.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var attributes interface{}
//   var metrics interface{}
//
//   campaignEventFilterProperty := &CampaignEventFilterProperty{
//   	Dimensions: &EventDimensionsProperty{
//   		Attributes: attributes,
//   		EventType: &SetDimensionProperty{
//   			DimensionType: jsii.String("dimensionType"),
//   			Values: []*string{
//   				jsii.String("values"),
//   			},
//   		},
//   		Metrics: metrics,
//   	},
//   	FilterType: jsii.String("filterType"),
//   }
//
type CfnCampaign_CampaignEventFilterProperty struct {
	// The dimension settings of the event filter for the campaign.
	Dimensions interface{} `field:"optional" json:"dimensions" yaml:"dimensions"`
	// The type of event that causes the campaign to be sent.
	//
	// Valid values are: `SYSTEM` , sends the campaign when a system event occurs; and, `ENDPOINT` , sends the campaign when an endpoint event (Events resource) occurs.
	FilterType *string `field:"optional" json:"filterType" yaml:"filterType"`
}

