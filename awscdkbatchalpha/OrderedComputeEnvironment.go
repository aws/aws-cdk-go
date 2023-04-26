// The CDK Construct Library for AWS::Batch
package awscdkbatchalpha


// Assigns an order to a ComputeEnvironment.
//
// The JobQueue will prioritize the lowest-order ComputeEnvironment.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import batch_alpha "github.com/aws/aws-cdk-go/awscdkbatchalpha"
//
//   var computeEnvironment iComputeEnvironment
//
//   orderedComputeEnvironment := &OrderedComputeEnvironment{
//   	ComputeEnvironment: computeEnvironment,
//   	Order: jsii.Number(123),
//   }
//
// Experimental.
type OrderedComputeEnvironment struct {
	// The ComputeEnvironment to link to this JobQueue.
	// Experimental.
	ComputeEnvironment IComputeEnvironment `field:"required" json:"computeEnvironment" yaml:"computeEnvironment"`
	// The order associated with `computeEnvironment`.
	// Experimental.
	Order *float64 `field:"required" json:"order" yaml:"order"`
}

