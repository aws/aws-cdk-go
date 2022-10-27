package awsbudgets


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   historicalOptionsProperty := &historicalOptionsProperty{
//   	budgetAdjustmentPeriod: jsii.Number(123),
//   }
//
type CfnBudget_HistoricalOptionsProperty struct {
	// `CfnBudget.HistoricalOptionsProperty.BudgetAdjustmentPeriod`.
	BudgetAdjustmentPeriod *float64 `field:"required" json:"budgetAdjustmentPeriod" yaml:"budgetAdjustmentPeriod"`
}

