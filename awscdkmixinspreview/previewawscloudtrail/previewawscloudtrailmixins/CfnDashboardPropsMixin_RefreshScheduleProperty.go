package previewawscloudtrailmixins


// The schedule for a dashboard refresh.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   refreshScheduleProperty := &RefreshScheduleProperty{
//   	Frequency: &FrequencyProperty{
//   		Unit: jsii.String("unit"),
//   		Value: jsii.Number(123),
//   	},
//   	Status: jsii.String("status"),
//   	TimeOfDay: jsii.String("timeOfDay"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudtrail-dashboard-refreshschedule.html
//
type CfnDashboardPropsMixin_RefreshScheduleProperty struct {
	// The frequency at which you want the dashboard refreshed.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudtrail-dashboard-refreshschedule.html#cfn-cloudtrail-dashboard-refreshschedule-frequency
	//
	Frequency interface{} `field:"optional" json:"frequency" yaml:"frequency"`
	// Specifies whether the refresh schedule is enabled.
	//
	// Set the value to `ENABLED` to enable the refresh schedule, or to `DISABLED` to turn off the refresh schedule.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudtrail-dashboard-refreshschedule.html#cfn-cloudtrail-dashboard-refreshschedule-status
	//
	Status *string `field:"optional" json:"status" yaml:"status"`
	// The time of day in UTC to run the schedule;
	//
	// for hourly only refer to minutes; default is 00:00.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudtrail-dashboard-refreshschedule.html#cfn-cloudtrail-dashboard-refreshschedule-timeofday
	//
	TimeOfDay *string `field:"optional" json:"timeOfDay" yaml:"timeOfDay"`
}

