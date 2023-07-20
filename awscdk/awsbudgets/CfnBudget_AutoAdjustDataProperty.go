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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-budgets-budget-autoadjustdata.html
//
type CfnBudget_AutoAdjustDataProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-budgets-budget-autoadjustdata.html#cfn-budgets-budget-autoadjustdata-autoadjusttype
	//
	AutoAdjustType *string `field:"required" json:"autoAdjustType" yaml:"autoAdjustType"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-budgets-budget-autoadjustdata.html#cfn-budgets-budget-autoadjustdata-historicaloptions
	//
	HistoricalOptions interface{} `field:"optional" json:"historicalOptions" yaml:"historicalOptions"`
}

