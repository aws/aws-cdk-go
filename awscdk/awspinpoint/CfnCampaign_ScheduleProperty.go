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
//   scheduleProperty := &ScheduleProperty{
//   	EndTime: jsii.String("endTime"),
//   	EventFilter: &CampaignEventFilterProperty{
//   		Dimensions: &EventDimensionsProperty{
//   			Attributes: attributes,
//   			EventType: &SetDimensionProperty{
//   				DimensionType: jsii.String("dimensionType"),
//   				Values: []*string{
//   					jsii.String("values"),
//   				},
//   			},
//   			Metrics: metrics,
//   		},
//   		FilterType: jsii.String("filterType"),
//   	},
//   	Frequency: jsii.String("frequency"),
//   	IsLocalTime: jsii.Boolean(false),
//   	QuietTime: &QuietTimeProperty{
//   		End: jsii.String("end"),
//   		Start: jsii.String("start"),
//   	},
//   	StartTime: jsii.String("startTime"),
//   	TimeZone: jsii.String("timeZone"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pinpoint-campaign-schedule.html
//
type CfnCampaign_ScheduleProperty struct {
	// The scheduled time, in ISO 8601 format, when the campaign ended or will end.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pinpoint-campaign-schedule.html#cfn-pinpoint-campaign-schedule-endtime
	//
	EndTime *string `field:"optional" json:"endTime" yaml:"endTime"`
	// The type of event that causes the campaign to be sent, if the value of the `Frequency` property is `EVENT` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pinpoint-campaign-schedule.html#cfn-pinpoint-campaign-schedule-eventfilter
	//
	EventFilter interface{} `field:"optional" json:"eventFilter" yaml:"eventFilter"`
	// Specifies how often the campaign is sent or whether the campaign is sent in response to a specific event.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pinpoint-campaign-schedule.html#cfn-pinpoint-campaign-schedule-frequency
	//
	Frequency *string `field:"optional" json:"frequency" yaml:"frequency"`
	// Specifies whether the start and end times for the campaign schedule use each recipient's local time.
	//
	// To base the schedule on each recipient's local time, set this value to `true` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pinpoint-campaign-schedule.html#cfn-pinpoint-campaign-schedule-islocaltime
	//
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
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pinpoint-campaign-schedule.html#cfn-pinpoint-campaign-schedule-quiettime
	//
	QuietTime interface{} `field:"optional" json:"quietTime" yaml:"quietTime"`
	// The scheduled time when the campaign began or will begin.
	//
	// Valid values are: `IMMEDIATE` , to start the campaign immediately; or, a specific time in ISO 8601 format.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pinpoint-campaign-schedule.html#cfn-pinpoint-campaign-schedule-starttime
	//
	StartTime *string `field:"optional" json:"startTime" yaml:"startTime"`
	// The starting UTC offset for the campaign schedule, if the value of the `IsLocalTime` property is `true` .
	//
	// Valid values are: `UTC, UTC+01, UTC+02, UTC+03, UTC+03:30, UTC+04, UTC+04:30, UTC+05, UTC+05:30, UTC+05:45, UTC+06, UTC+06:30, UTC+07, UTC+08, UTC+09, UTC+09:30, UTC+10, UTC+10:30, UTC+11, UTC+12, UTC+13, UTC-02, UTC-03, UTC-04, UTC-05, UTC-06, UTC-07, UTC-08, UTC-09, UTC-10,` and `UTC-11` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pinpoint-campaign-schedule.html#cfn-pinpoint-campaign-schedule-timezone
	//
	TimeZone *string `field:"optional" json:"timeZone" yaml:"timeZone"`
}

