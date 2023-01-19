package awsbudgets


// The parameters that determine the budget amount for an auto-adjusting budget.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   autoAdjustDataProperty := &autoAdjustDataProperty{
//   	autoAdjustType: jsii.String("autoAdjustType"),
//
//   	// the properties below are optional
//   	historicalOptions: &historicalOptionsProperty{
//   		budgetAdjustmentPeriod: jsii.Number(123),
//   	},
//   }
//
type CfnBudget_AutoAdjustDataProperty struct {
	// The string that defines whether your budget auto-adjusts based on historical or forecasted data.
	AutoAdjustType *string `field:"required" json:"autoAdjustType" yaml:"autoAdjustType"`
	// The parameters that define or describe the historical data that your auto-adjusting budget is based on.
	HistoricalOptions interface{} `field:"optional" json:"historicalOptions" yaml:"historicalOptions"`
}

