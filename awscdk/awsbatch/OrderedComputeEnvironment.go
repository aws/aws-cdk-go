package awsbatch

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsbatch"
)

// Assigns an order to a ComputeEnvironment.
//
// The JobQueue will prioritize the lowest-order ComputeEnvironment.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var computeEnvironmentRef IComputeEnvironmentRef
//
//   orderedComputeEnvironment := &OrderedComputeEnvironment{
//   	ComputeEnvironment: computeEnvironmentRef,
//   	Order: jsii.Number(123),
//   }
//
type OrderedComputeEnvironment struct {
	// The ComputeEnvironment to link to this JobQueue.
	ComputeEnvironment interfacesawsbatch.IComputeEnvironmentRef `field:"required" json:"computeEnvironment" yaml:"computeEnvironment"`
	// The order associated with `computeEnvironment`.
	Order *float64 `field:"required" json:"order" yaml:"order"`
}

