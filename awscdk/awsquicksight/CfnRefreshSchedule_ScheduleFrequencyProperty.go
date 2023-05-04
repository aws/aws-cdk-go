package awsquicksight


// The frequency for the refresh schedule.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   scheduleFrequencyProperty := &ScheduleFrequencyProperty{
//   	Interval: jsii.String("interval"),
//   	RefreshOnDay: &RefreshOnDayProperty{
//   		DayOfMonth: jsii.String("dayOfMonth"),
//   		DayOfWeek: jsii.String("dayOfWeek"),
//   	},
//   	TimeOfTheDay: jsii.String("timeOfTheDay"),
//   	TimeZone: jsii.String("timeZone"),
//   }
//
type CfnRefreshSchedule_ScheduleFrequencyProperty struct {
	// The interval between scheduled refreshes. Valid values are as follows:.
	//
	// - `MINUTE15` : The dataset refreshes every 15 minutes. This value is only supported for incremental refreshes. This interval can only be used for one schedule per dataset.
	// - `MINUTE30` : The dataset refreshes every 30 minutes. This value is only supported for incremental refreshes. This interval can only be used for one schedule per dataset.
	// - `HOURLY` : The dataset refreshes every hour. This interval can only be used for one schedule per dataset.
	// - `DAILY` : The dataset refreshes every day.
	// - `WEEKLY` : The dataset refreshes every week.
	// - `MONTHLY` : The dataset refreshes every month.
	Interval *string `field:"optional" json:"interval" yaml:"interval"`
	// The day of the week that you want to schedule the refresh on.
	//
	// This value is required for weekly and monthly refresh intervals.
	RefreshOnDay interface{} `field:"optional" json:"refreshOnDay" yaml:"refreshOnDay"`
	// The time of day that you want the dataset to refresh.
	//
	// This value is expressed in HH:MM format. This field is not required for schedules that refresh hourly.
	TimeOfTheDay *string `field:"optional" json:"timeOfTheDay" yaml:"timeOfTheDay"`
	// The timezone that you want the refresh schedule to use.
	//
	// The timezone ID must match a corresponding ID found on `java.util.time.getAvailableIDs()` .
	TimeZone *string `field:"optional" json:"timeZone" yaml:"timeZone"`
}

