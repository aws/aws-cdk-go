package awsbudgets


// The parameters that define or describe the historical data that your auto-adjusting budget is based on.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   historicalOptionsProperty := &HistoricalOptionsProperty{
//   	BudgetAdjustmentPeriod: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-budgets-budget-historicaloptions.html
//
type CfnBudget_HistoricalOptionsProperty struct {
	// The number of budget periods included in the moving-average calculation that determines your auto-adjusted budget amount.
	//
	// The maximum value depends on the `TimeUnit` granularity of the budget:
	//
	// - For the `DAILY` granularity, the maximum value is `60` .
	// - For the `MONTHLY` granularity, the maximum value is `12` .
	// - For the `QUARTERLY` granularity, the maximum value is `4` .
	// - For the `ANNUALLY` granularity, the maximum value is `1` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-budgets-budget-historicaloptions.html#cfn-budgets-budget-historicaloptions-budgetadjustmentperiod
	//
	BudgetAdjustmentPeriod *float64 `field:"required" json:"budgetAdjustmentPeriod" yaml:"budgetAdjustmentPeriod"`
}

