package awsbudgets


// Determine the budget amount for an auto-adjusting budget.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   autoAdjustDataProperty := &AutoAdjustDataProperty{
//   	AutoAdjustType: jsii.String("autoAdjustType"),
//
//   	// the properties below are optional
//   	HistoricalOptions: &HistoricalOptionsProperty{
//   		BudgetAdjustmentPeriod: jsii.Number(123),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-budgets-budget-autoadjustdata.html
//
type CfnBudget_AutoAdjustDataProperty struct {
	// The string that defines whether your budget auto-adjusts based on historical or forecasted data.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-budgets-budget-autoadjustdata.html#cfn-budgets-budget-autoadjustdata-autoadjusttype
	//
	AutoAdjustType *string `field:"required" json:"autoAdjustType" yaml:"autoAdjustType"`
	// The parameters that define or describe the historical data that your auto-adjusting budget is based on.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-budgets-budget-autoadjustdata.html#cfn-budgets-budget-autoadjustdata-historicaloptions
	//
	HistoricalOptions interface{} `field:"optional" json:"historicalOptions" yaml:"historicalOptions"`
}

