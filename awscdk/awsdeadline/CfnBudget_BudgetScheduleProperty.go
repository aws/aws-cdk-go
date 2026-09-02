package awsdeadline


// The start and end time of the budget.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   budgetScheduleProperty := &BudgetScheduleProperty{
//   	Fixed: &FixedBudgetScheduleProperty{
//   		EndTime: jsii.String("endTime"),
//   		StartTime: jsii.String("startTime"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-deadline-budget-budgetschedule.html
//
type CfnBudget_BudgetScheduleProperty struct {
	// The details of a fixed budget schedule.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-deadline-budget-budgetschedule.html#cfn-deadline-budget-budgetschedule-fixed
	//
	Fixed interface{} `field:"required" json:"fixed" yaml:"fixed"`
}

