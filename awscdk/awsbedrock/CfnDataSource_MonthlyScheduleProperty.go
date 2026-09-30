package awsbedrock


// A monthly refresh on a specified day of the month.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var lastDayOfMonth interface{}
//
//   monthlyScheduleProperty := &MonthlyScheduleProperty{
//   	DayOfMonth: &DayOfMonthProperty{
//   		DayNumber: jsii.Number(123),
//   		LastDayOfMonth: lastDayOfMonth,
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-datasource-monthlyschedule.html
//
type CfnDataSource_MonthlyScheduleProperty struct {
	// Day of the month on which a monthly refresh runs.
	//
	// Exactly one variant is set: an explicit day number, or the last calendar day of the month.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-datasource-monthlyschedule.html#cfn-bedrock-datasource-monthlyschedule-dayofmonth
	//
	DayOfMonth interface{} `field:"required" json:"dayOfMonth" yaml:"dayOfMonth"`
}

