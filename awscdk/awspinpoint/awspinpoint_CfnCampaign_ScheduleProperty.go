package awspinpoint


// Specifies the schedule settings for a campaign.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var attributes interface{}
//   var metrics interface{}
//
//   scheduleProperty := &scheduleProperty{
//   	endTime: jsii.String("endTime"),
//   	eventFilter: &campaignEventFilterProperty{
//   		dimensions: &eventDimensionsProperty{
//   			attributes: attributes,
//   			eventType: &setDimensionProperty{
//   				dimensionType: jsii.String("dimensionType"),
//   				values: []*string{
//   					jsii.String("values"),
//   				},
//   			},
//   			metrics: metrics,
//   		},
//   		filterType: jsii.String("filterType"),
//   	},
//   	frequency: jsii.String("frequency"),
//   	isLocalTime: jsii.Boolean(false),
//   	quietTime: &quietTimeProperty{
//   		end: jsii.String("end"),
//   		start: jsii.String("start"),
//   	},
//   	startTime: jsii.String("startTime"),
//   	timeZone: jsii.String("timeZone"),
//   }
//
type CfnCampaign_ScheduleProperty struct {
	// The scheduled time, in ISO 8601 format, when the campaign ended or will end.
	EndTime *string `field:"optional" json:"endTime" yaml:"endTime"`
	// The type of event that causes the campaign to be sent, if the value of the `Frequency` property is `EVENT` .
	EventFilter interface{} `field:"optional" json:"eventFilter" yaml:"eventFilter"`
	// Specifies how often the campaign is sent or whether the campaign is sent in response to a specific event.
	Frequency *string `field:"optional" json:"frequency" yaml:"frequency"`
	// Specifies whether the start and end times for the campaign schedule use each recipient's local time.
	//
	// To base the schedule on each recipient's local time, set this value to `true` .
	IsLocalTime interface{} `field:"optional" json:"isLocalTime" yaml:"isLocalTime"`
	// The default quiet time for the campaign.
	//
	// Quiet time is a specific time range when a campaign doesn't send messages to endpoints, if all the following conditions are met:
	//
	// - The `EndpointDemographic.Timezone` property of the endpoint is set to a valid value.
	// - The current time in the endpoint's time zone is later than or equal to the time specified by the `QuietTime.Start` property for the campaign.
	// - The current time in the endpoint's time zone is earlier than or equal to the time specified by the `QuietTime.End` property for the campaign.
	//
	// If any of the preceding conditions isn't met, the endpoint will receive messages from the campaign, even if quiet time is enabled.
	QuietTime interface{} `field:"optional" json:"quietTime" yaml:"quietTime"`
	// The scheduled time when the campaign began or will begin.
	//
	// Valid values are: `IMMEDIATE` , to start the campaign immediately; or, a specific time in ISO 8601 format.
	StartTime *string `field:"optional" json:"startTime" yaml:"startTime"`
	// The starting UTC offset for the campaign schedule, if the value of the `IsLocalTime` property is `true` .
	//
	// Valid values are: `UTC, UTC+01, UTC+02, UTC+03, UTC+03:30, UTC+04, UTC+04:30, UTC+05, UTC+05:30, UTC+05:45, UTC+06, UTC+06:30, UTC+07, UTC+08, UTC+09, UTC+09:30, UTC+10, UTC+10:30, UTC+11, UTC+12, UTC+13, UTC-02, UTC-03, UTC-04, UTC-05, UTC-06, UTC-07, UTC-08, UTC-09, UTC-10,` and `UTC-11` .
	TimeZone *string `field:"optional" json:"timeZone" yaml:"timeZone"`
}

