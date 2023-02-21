package awsbudgets


// The amount of cost or usage that's measured for a budget.
//
// For example, a `Spend` for `3 GB` of S3 usage has the following parameters:
//
// - An `Amount` of `3`
// - A `unit` of `GB`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   spendProperty := &spendProperty{
//   	amount: jsii.Number(123),
//   	unit: jsii.String("unit"),
//   }
//
type CfnBudget_SpendProperty struct {
	// The cost or usage amount that's associated with a budget forecast, actual spend, or budget threshold.
	Amount *float64 `field:"required" json:"amount" yaml:"amount"`
	// The unit of measurement that's used for the budget forecast, actual spend, or budget threshold, such as USD or GBP.
	Unit *string `field:"required" json:"unit" yaml:"unit"`
}

