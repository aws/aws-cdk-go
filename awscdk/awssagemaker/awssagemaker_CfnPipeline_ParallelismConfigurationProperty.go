package awssagemaker


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   parallelismConfigurationProperty := &parallelismConfigurationProperty{
//   	maxParallelExecutionSteps: jsii.Number(123),
//   }
//
type CfnPipeline_ParallelismConfigurationProperty struct {
	// `CfnPipeline.ParallelismConfigurationProperty.MaxParallelExecutionSteps`.
	MaxParallelExecutionSteps *float64 `field:"required" json:"maxParallelExecutionSteps" yaml:"maxParallelExecutionSteps"`
}

