package interfacesawsbudgets


// A reference to a BudgetsAction resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   budgetsActionReference := &BudgetsActionReference{
//   	ActionId: jsii.String("actionId"),
//   	BudgetName: jsii.String("budgetName"),
//   }
//
type BudgetsActionReference struct {
	// The ActionId of the BudgetsAction resource.
	ActionId *string `field:"required" json:"actionId" yaml:"actionId"`
	// The BudgetName of the BudgetsAction resource.
	BudgetName *string `field:"required" json:"budgetName" yaml:"budgetName"`
}

