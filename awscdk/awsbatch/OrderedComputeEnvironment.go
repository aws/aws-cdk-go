package awsbatch


// Assigns an order to a ComputeEnvironment.
//
// The JobQueue will prioritize the lowest-order ComputeEnvironment.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var computeEnvironment IComputeEnvironment
//
//   orderedComputeEnvironment := &OrderedComputeEnvironment{
//   	ComputeEnvironment: computeEnvironment,
//   	Order: jsii.Number(123),
//   }
//
type OrderedComputeEnvironment struct {
	// The ComputeEnvironment to link to this JobQueue.
	ComputeEnvironment IComputeEnvironment `field:"required" json:"computeEnvironment" yaml:"computeEnvironment"`
	// The order associated with `computeEnvironment`.
	Order *float64 `field:"required" json:"order" yaml:"order"`
}

