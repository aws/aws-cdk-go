package awssagemaker


// Configuration that controls the parallelism of the pipeline.
//
// By default, the parallelism configuration specified applies to all executions of the pipeline unless overridden.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   parallelismConfigurationProperty := &ParallelismConfigurationProperty{
//   	MaxParallelExecutionSteps: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-pipeline-parallelismconfiguration.html
//
type CfnPipeline_ParallelismConfigurationProperty struct {
	// The max number of steps that can be executed in parallel.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-pipeline-parallelismconfiguration.html#cfn-sagemaker-pipeline-parallelismconfiguration-maxparallelexecutionsteps
	//
	MaxParallelExecutionSteps *float64 `field:"required" json:"maxParallelExecutionSteps" yaml:"maxParallelExecutionSteps"`
}

