package awsbudgets


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
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-budgets-budget-historicaloptions.html#cfn-budgets-budget-historicaloptions-budgetadjustmentperiod
	//
	BudgetAdjustmentPeriod *float64 `field:"required" json:"budgetAdjustmentPeriod" yaml:"budgetAdjustmentPeriod"`
}

