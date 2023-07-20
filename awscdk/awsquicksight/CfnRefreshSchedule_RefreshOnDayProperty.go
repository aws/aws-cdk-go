package awsquicksight


// The day that you want yout dataset to refresh.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   refreshOnDayProperty := &RefreshOnDayProperty{
//   	DayOfMonth: jsii.String("dayOfMonth"),
//   	DayOfWeek: jsii.String("dayOfWeek"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-refreshschedule-refreshonday.html
//
type CfnRefreshSchedule_RefreshOnDayProperty struct {
	// The day of the month that you want your dataset to refresh.
	//
	// This value is required for monthly refresh intervals.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-refreshschedule-refreshonday.html#cfn-quicksight-refreshschedule-refreshonday-dayofmonth
	//
	DayOfMonth *string `field:"optional" json:"dayOfMonth" yaml:"dayOfMonth"`
	// The day of the week that you want to schedule the refresh on.
	//
	// This value is required for weekly and monthly refresh intervals.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-quicksight-refreshschedule-refreshonday.html#cfn-quicksight-refreshschedule-refreshonday-dayofweek
	//
	DayOfWeek *string `field:"optional" json:"dayOfWeek" yaml:"dayOfWeek"`
}

