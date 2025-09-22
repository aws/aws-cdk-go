package awsce


// A reference to a CostCategory resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   costCategoryReference := &CostCategoryReference{
//   	CostCategoryArn: jsii.String("costCategoryArn"),
//   }
//
type CostCategoryReference struct {
	// The Arn of the CostCategory resource.
	CostCategoryArn *string `field:"required" json:"costCategoryArn" yaml:"costCategoryArn"`
}

