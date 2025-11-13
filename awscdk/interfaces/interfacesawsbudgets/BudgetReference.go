package interfacesawsbudgets


// A reference to a Budget resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   budgetReference := &BudgetReference{
//   	BudgetId: jsii.String("budgetId"),
//   }
//
type BudgetReference struct {
	// The Id of the Budget resource.
	BudgetId *string `field:"required" json:"budgetId" yaml:"budgetId"`
}

