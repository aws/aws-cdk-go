package awsdeadline


// The budget action to add.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   budgetActionToAddProperty := &BudgetActionToAddProperty{
//   	Description: jsii.String("description"),
//   	ThresholdPercentage: jsii.Number(123),
//   	Type: jsii.String("type"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-deadline-budget-budgetactiontoadd.html
//
type CfnBudgetPropsMixin_BudgetActionToAddProperty struct {
	// A description for the budget action.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-deadline-budget-budgetactiontoadd.html#cfn-deadline-budget-budgetactiontoadd-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The percentage threshold for the budget action.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-deadline-budget-budgetactiontoadd.html#cfn-deadline-budget-budgetactiontoadd-thresholdpercentage
	//
	ThresholdPercentage *float64 `field:"optional" json:"thresholdPercentage" yaml:"thresholdPercentage"`
	// The type of budget action.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-deadline-budget-budgetactiontoadd.html#cfn-deadline-budget-budgetactiontoadd-type
	//
	Type *string `field:"optional" json:"type" yaml:"type"`
}

