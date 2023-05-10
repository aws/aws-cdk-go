package awsbudgets


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
type CfnBudget_AutoAdjustDataProperty struct {
	// `CfnBudget.AutoAdjustDataProperty.AutoAdjustType`.
	AutoAdjustType *string `field:"required" json:"autoAdjustType" yaml:"autoAdjustType"`
	// `CfnBudget.AutoAdjustDataProperty.HistoricalOptions`.
	HistoricalOptions interface{} `field:"optional" json:"historicalOptions" yaml:"historicalOptions"`
}

