package awsbedrock


// Recurring schedule on which the connector automatically refreshes ingested content.
//
// Exactly one frequency variant is set.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var daily interface{}
//   var lastDayOfMonth interface{}
//
//   syncScheduleProperty := &SyncScheduleProperty{
//   	Daily: daily,
//   	Monthly: &MonthlyScheduleProperty{
//   		DayOfMonth: &DayOfMonthProperty{
//   			DayNumber: jsii.Number(123),
//   			LastDayOfMonth: lastDayOfMonth,
//   		},
//   	},
//   	Weekly: &WeeklyScheduleProperty{
//   		DayOfWeek: jsii.String("dayOfWeek"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-datasource-syncschedule.html
//
type CfnDataSource_SyncScheduleProperty struct {
	// A daily refresh.
	//
	// The run time is system-chosen (off-peak) and not customer-configurable.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-datasource-syncschedule.html#cfn-bedrock-datasource-syncschedule-daily
	//
	Daily interface{} `field:"optional" json:"daily" yaml:"daily"`
	// A monthly refresh on a specified day of the month.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-datasource-syncschedule.html#cfn-bedrock-datasource-syncschedule-monthly
	//
	Monthly interface{} `field:"optional" json:"monthly" yaml:"monthly"`
	// A weekly refresh on a specified day of the week.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-datasource-syncschedule.html#cfn-bedrock-datasource-syncschedule-weekly
	//
	Weekly interface{} `field:"optional" json:"weekly" yaml:"weekly"`
}

