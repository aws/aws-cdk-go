package awsbudgets


// The period of time that is covered by a budget.
//
// The period has a start date and an end date. The start date must come before the end date. There are no restrictions on the end date.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   timePeriodProperty := &timePeriodProperty{
//   	end: jsii.String("end"),
//   	start: jsii.String("start"),
//   }
//
type CfnBudget_TimePeriodProperty struct {
	// The end date for a budget.
	//
	// If you didn't specify an end date, AWS set your end date to `06/15/87 00:00 UTC` . The defaults are the same for the AWS Billing and Cost Management console and the API.
	//
	// After the end date, AWS deletes the budget and all the associated notifications and subscribers. You can change your end date with the `UpdateBudget` operation.
	End *string `field:"optional" json:"end" yaml:"end"`
	// The start date for a budget.
	//
	// If you created your budget and didn't specify a start date, the start date defaults to the start of the chosen time period (MONTHLY, QUARTERLY, or ANNUALLY). For example, if you create your budget on January 24, 2019, choose `MONTHLY` , and don't set a start date, the start date defaults to `01/01/19 00:00 UTC` . The defaults are the same for the AWS Billing and Cost Management console and the API.
	//
	// You can change your start date with the `UpdateBudget` operation.
	//
	// Valid values depend on the value of `BudgetType` :
	//
	// - If `BudgetType` is `COST` or `USAGE` : Valid values are `MONTHLY` , `QUARTERLY` , and `ANNUALLY` .
	// - If `BudgetType` is `RI_UTILIZATION` or `RI_COVERAGE` : Valid values are `DAILY` , `MONTHLY` , `QUARTERLY` , and `ANNUALLY` .
	Start *string `field:"optional" json:"start" yaml:"start"`
}

