package awsglue


// A reference to a TableOptimizer resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   tableOptimizerReference := &TableOptimizerReference{
//   	TableOptimizerId: jsii.String("tableOptimizerId"),
//   }
//
type TableOptimizerReference struct {
	// The Id of the TableOptimizer resource.
	TableOptimizerId *string `field:"required" json:"tableOptimizerId" yaml:"tableOptimizerId"`
}

