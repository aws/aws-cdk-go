package awsbudgets


// The types of cost that are included in a `COST` budget, such as tax and subscriptions.
//
// `USAGE` , `RI_UTILIZATION` , `RI_COVERAGE` , `SAVINGS_PLANS_UTILIZATION` , and `SAVINGS_PLANS_COVERAGE` budgets don't have `CostTypes` .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   costTypesProperty := &CostTypesProperty{
//   	IncludeCredit: jsii.Boolean(false),
//   	IncludeDiscount: jsii.Boolean(false),
//   	IncludeOtherSubscription: jsii.Boolean(false),
//   	IncludeRecurring: jsii.Boolean(false),
//   	IncludeRefund: jsii.Boolean(false),
//   	IncludeSubscription: jsii.Boolean(false),
//   	IncludeSupport: jsii.Boolean(false),
//   	IncludeTax: jsii.Boolean(false),
//   	IncludeUpfront: jsii.Boolean(false),
//   	UseAmortized: jsii.Boolean(false),
//   	UseBlended: jsii.Boolean(false),
//   }
//
type CfnBudget_CostTypesProperty struct {
	// Specifies whether a budget includes credits.
	//
	// The default value is `true` .
	IncludeCredit interface{} `field:"optional" json:"includeCredit" yaml:"includeCredit"`
	// Specifies whether a budget includes discounts.
	//
	// The default value is `true` .
	IncludeDiscount interface{} `field:"optional" json:"includeDiscount" yaml:"includeDiscount"`
	// Specifies whether a budget includes non-RI subscription costs.
	//
	// The default value is `true` .
	IncludeOtherSubscription interface{} `field:"optional" json:"includeOtherSubscription" yaml:"includeOtherSubscription"`
	// Specifies whether a budget includes recurring fees such as monthly RI fees.
	//
	// The default value is `true` .
	IncludeRecurring interface{} `field:"optional" json:"includeRecurring" yaml:"includeRecurring"`
	// Specifies whether a budget includes refunds.
	//
	// The default value is `true` .
	IncludeRefund interface{} `field:"optional" json:"includeRefund" yaml:"includeRefund"`
	// Specifies whether a budget includes subscriptions.
	//
	// The default value is `true` .
	IncludeSubscription interface{} `field:"optional" json:"includeSubscription" yaml:"includeSubscription"`
	// Specifies whether a budget includes support subscription fees.
	//
	// The default value is `true` .
	IncludeSupport interface{} `field:"optional" json:"includeSupport" yaml:"includeSupport"`
	// Specifies whether a budget includes taxes.
	//
	// The default value is `true` .
	IncludeTax interface{} `field:"optional" json:"includeTax" yaml:"includeTax"`
	// Specifies whether a budget includes upfront RI costs.
	//
	// The default value is `true` .
	IncludeUpfront interface{} `field:"optional" json:"includeUpfront" yaml:"includeUpfront"`
	// Specifies whether a budget uses the amortized rate.
	//
	// The default value is `false` .
	UseAmortized interface{} `field:"optional" json:"useAmortized" yaml:"useAmortized"`
	// Specifies whether a budget uses a blended rate.
	//
	// The default value is `false` .
	UseBlended interface{} `field:"optional" json:"useBlended" yaml:"useBlended"`
}

