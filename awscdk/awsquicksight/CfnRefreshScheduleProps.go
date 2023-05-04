package awsquicksight


// Properties for defining a `CfnRefreshSchedule`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnRefreshScheduleProps := &CfnRefreshScheduleProps{
//   	AwsAccountId: jsii.String("awsAccountId"),
//   	DataSetId: jsii.String("dataSetId"),
//   	Schedule: &RefreshScheduleMapProperty{
//   		RefreshType: jsii.String("refreshType"),
//   		ScheduleFrequency: &ScheduleFrequencyProperty{
//   			Interval: jsii.String("interval"),
//   			RefreshOnDay: &RefreshOnDayProperty{
//   				DayOfMonth: jsii.String("dayOfMonth"),
//   				DayOfWeek: jsii.String("dayOfWeek"),
//   			},
//   			TimeOfTheDay: jsii.String("timeOfTheDay"),
//   			TimeZone: jsii.String("timeZone"),
//   		},
//   		ScheduleId: jsii.String("scheduleId"),
//   		StartAfterDateTime: jsii.String("startAfterDateTime"),
//   	},
//   }
//
type CfnRefreshScheduleProps struct {
	// The AWS account ID of the account that you are creating a schedule in.
	AwsAccountId *string `field:"optional" json:"awsAccountId" yaml:"awsAccountId"`
	// The ID of the dataset that you are creating a refresh schedule for.
	DataSetId *string `field:"optional" json:"dataSetId" yaml:"dataSetId"`
	// The refresh schedule of a dataset.
	Schedule interface{} `field:"optional" json:"schedule" yaml:"schedule"`
}

