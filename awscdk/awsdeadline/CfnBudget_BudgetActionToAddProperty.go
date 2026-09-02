package awsdeadline


// The budget action to add.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   budgetActionToAddProperty := &BudgetActionToAddProperty{
//   	ThresholdPercentage: jsii.Number(123),
//   	Type: jsii.String("type"),
//
//   	// the properties below are optional
//   	Description: jsii.String("description"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-deadline-budget-budgetactiontoadd.html
//
type CfnBudget_BudgetActionToAddProperty struct {
	// The percentage threshold for the budget action.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-deadline-budget-budgetactiontoadd.html#cfn-deadline-budget-budgetactiontoadd-thresholdpercentage
	//
	ThresholdPercentage *float64 `field:"required" json:"thresholdPercentage" yaml:"thresholdPercentage"`
	// The type of budget action.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-deadline-budget-budgetactiontoadd.html#cfn-deadline-budget-budgetactiontoadd-type
	//
	Type *string `field:"required" json:"type" yaml:"type"`
	// A description for the budget action.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-deadline-budget-budgetactiontoadd.html#cfn-deadline-budget-budgetactiontoadd-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
}

