// The CDK Construct Library for AWS::Batch
package awscdkbatchalpha


// Properties for mapping a compute environment to a job queue.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import batch_alpha "github.com/aws/aws-cdk-go/awscdkbatchalpha"
//
//   var computeEnvironment computeEnvironment
//
//   jobQueueComputeEnvironment := &jobQueueComputeEnvironment{
//   	computeEnvironment: computeEnvironment,
//   	order: jsii.Number(123),
//   }
//
// Experimental.
type JobQueueComputeEnvironment struct {
	// The batch compute environment to use for processing submitted jobs to this queue.
	// Experimental.
	ComputeEnvironment IComputeEnvironment `field:"required" json:"computeEnvironment" yaml:"computeEnvironment"`
	// The order in which this compute environment will be selected for dynamic allocation of resources to process submitted jobs.
	// Experimental.
	Order *float64 `field:"required" json:"order" yaml:"order"`
}

