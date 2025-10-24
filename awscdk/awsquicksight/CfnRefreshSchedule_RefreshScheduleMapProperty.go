package awsquicksight


// A summary of a configured refresh schedule for a dataset.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   refreshScheduleMapProperty := &RefreshScheduleMapProperty{
//   	RefreshType: jsii.String("refreshType"),
//   	ScheduleFrequency: &ScheduleFrequencyProperty{
//   		Interval: jsii.String("interval"),
//   		RefreshOnDay: &RefreshOnDayProperty{
//   			DayOfMonth: jsii.String("dayOfMonth"),
//   			DayOfWeek: jsii.String("dayOfWeek"),
//   		},
//   		TimeOfTheDay: jsii.String("timeOfTheDay"),
//   		TimeZone: jsii.String("timeZone"),
//   	},
//   	ScheduleId: jsii.String("scheduleId"),
//   	StartAfterDateTime: jsii.String("startAfterDateTime"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-refreshschedule-refreshschedulemap.html
//
type CfnRefreshSchedule_RefreshScheduleMapProperty struct {
	// The type of refresh that a dataset undergoes. Valid values are as follows:.
	//
	// - `FULL_REFRESH` : A complete refresh of a dataset.
	// - `INCREMENTAL_REFRESH` : A partial refresh of some rows of a dataset, based on the time window specified.
	//
	// For more information on full and incremental refreshes, see [Refreshing SPICE data](https://docs.aws.amazon.com/quicksight/latest/user/refreshing-imported-data.html) in the *Quick Suite User Guide* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-refreshschedule-refreshschedulemap.html#cfn-quicksight-refreshschedule-refreshschedulemap-refreshtype
	//
	RefreshType *string `field:"optional" json:"refreshType" yaml:"refreshType"`
	// The frequency for the refresh schedule.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-refreshschedule-refreshschedulemap.html#cfn-quicksight-refreshschedule-refreshschedulemap-schedulefrequency
	//
	ScheduleFrequency interface{} `field:"optional" json:"scheduleFrequency" yaml:"scheduleFrequency"`
	// An identifier for the refresh schedule.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-refreshschedule-refreshschedulemap.html#cfn-quicksight-refreshschedule-refreshschedulemap-scheduleid
	//
	ScheduleId *string `field:"optional" json:"scheduleId" yaml:"scheduleId"`
	// Time after which the refresh schedule can be started, expressed in `YYYY-MM-DDTHH:MM:SS` format.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-refreshschedule-refreshschedulemap.html#cfn-quicksight-refreshschedule-refreshschedulemap-startafterdatetime
	//
	StartAfterDateTime *string `field:"optional" json:"startAfterDateTime" yaml:"startAfterDateTime"`
}

