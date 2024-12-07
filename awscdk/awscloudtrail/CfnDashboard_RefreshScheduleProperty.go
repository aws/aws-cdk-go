package awscloudtrail


// Configures the automatic refresh schedule for the dashboard.
//
// Includes the frequency unit (DAYS or HOURS) and value, as well as the status (ENABLED or DISABLED) of the refresh schedule.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
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
type CfnDashboard_RefreshScheduleProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudtrail-dashboard-refreshschedule.html#cfn-cloudtrail-dashboard-refreshschedule-frequency
	//
	Frequency interface{} `field:"optional" json:"frequency" yaml:"frequency"`
	// The status of the schedule.
	//
	// Supported values are ENABLED and DISABLED.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudtrail-dashboard-refreshschedule.html#cfn-cloudtrail-dashboard-refreshschedule-status
	//
	Status *string `field:"optional" json:"status" yaml:"status"`
	// StartTime of the automatic schedule refresh.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudtrail-dashboard-refreshschedule.html#cfn-cloudtrail-dashboard-refreshschedule-timeofday
	//
	TimeOfDay *string `field:"optional" json:"timeOfDay" yaml:"timeOfDay"`
}

