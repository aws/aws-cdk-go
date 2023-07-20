package awsbatch


// The order that compute environments are tried in for job placement within a queue.
//
// Compute environments are tried in ascending order. For example, if two compute environments are associated with a job queue, the compute environment with a lower order integer value is tried for job placement first. Compute environments must be in the `VALID` state before you can associate them with a job queue. All of the compute environments must be either EC2 ( `EC2` or `SPOT` ) or Fargate ( `FARGATE` or `FARGATE_SPOT` ); EC2 and Fargate compute environments can't be mixed.
//
// > All compute environments that are associated with a job queue must share the same architecture. AWS Batch doesn't support mixing compute environment architecture types in a single job queue.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   computeEnvironmentOrderProperty := &ComputeEnvironmentOrderProperty{
//   	ComputeEnvironment: jsii.String("computeEnvironment"),
//   	Order: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-batch-jobqueue-computeenvironmentorder.html
//
type CfnJobQueue_ComputeEnvironmentOrderProperty struct {
	// The Amazon Resource Name (ARN) of the compute environment.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-batch-jobqueue-computeenvironmentorder.html#cfn-batch-jobqueue-computeenvironmentorder-computeenvironment
	//
	ComputeEnvironment *string `field:"required" json:"computeEnvironment" yaml:"computeEnvironment"`
	// The order of the compute environment.
	//
	// Compute environments are tried in ascending order. For example, if two compute environments are associated with a job queue, the compute environment with a lower `order` integer value is tried for job placement first.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-batch-jobqueue-computeenvironmentorder.html#cfn-batch-jobqueue-computeenvironmentorder-order
	//
	Order *float64 `field:"required" json:"order" yaml:"order"`
}

