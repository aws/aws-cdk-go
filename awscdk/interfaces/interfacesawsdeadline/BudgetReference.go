package interfacesawsdeadline


// A reference to a Budget resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   budgetReference := &BudgetReference{
//   	BudgetArn: jsii.String("budgetArn"),
//   }
//
type BudgetReference struct {
	// The Arn of the Budget resource.
	BudgetArn *string `field:"required" json:"budgetArn" yaml:"budgetArn"`
}

